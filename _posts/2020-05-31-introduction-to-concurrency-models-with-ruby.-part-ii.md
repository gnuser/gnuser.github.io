---
key: introduction-to-concurrency-models-with-ruby.-part-ii
title: Introduction to Concurrency Models with Ruby. Part II
date: 2020-05-31 15:31:36 +0800
tags: ruby GIL concurrency parallel event-machine
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

在本系列的第二部分中，我们将研究更高级的并发模型，如actor、通信顺序进程、软件事务内存，当然还有Guilds——一种可能在Ruby 3中实现的新并发模型。

<!--more-->

如果你没有读[第一部分](https://gnuser.github.io/2020/05/31/introduction-to-concurrency-models-with-ruby.-part-i.html#缺点), 建议先读一下。在那篇文章中，我描述了进程、线程、GIL、EventMachine和纤程，我将在本文中提到它们。

![img](/../../../../../../../media/2020-05-31-introduction-to-concurrency-models-with-ruby.-part-ii/1*wTtV8K5ZTTA8B0jv8aZqnQ.jpeg)

## Actors

Actor是并发原语，它可以相互发送消息、创建新的actor并确定如何响应下一个接收到的消息。它们保持自己的私有状态而不共享它，因此它们只能通过消息相互影响。因为没有共享状态，所以不需要锁。

> 不要通过共享内存进行通信;相反，通过通信共享内存。

Erlang和Scala都在语言本身中实现了Actor模型。在Ruby中，`Celluloid`是最流行的实现之一。在底层，它在一个单独的线程中运行每个Actor，并为每个方法调用使用纤程(Fiber)，以避免在等待其他Actor响应时阻塞方法。

这里有个使用 Celluloid 的例子：

```ruby
# actors.rb
require 'celluloid'
class Universe
  include Celluloid
  def say(msg)
    puts msg
    Celluloid::Actor[:world].say("#{msg} World!")
  end
end
class World
  include Celluloid
  def say(msg)
    puts msg
  end
end
Celluloid::Actor[:world] = World.new
Universe.new.say("Hello")
$ ruby actors.rb
Hello
Hello World!
```

### 优点：

- 没有多线程编程和共享内存意味着几乎无死锁，同步不用显式锁。
- 与Erlang类似，Celluloid让Actors具有容错能力，这意味着它会尝试用 [Supervisors](https://github.com/celluloid/celluloid/wiki/Supervisors)重启崩溃的Actors.
- Actor模型是为解决分布式程序的问题而设计的，所以它非常适合跨多台机器扩展。

### 缺点：

- 如果系统需要使用共享状态，或者您需要保证需要以特定顺序发生的行为，则Actors可能无法工作。
- 调试可能会很棘手——想象一下通过多个Actors跟踪系统流程，或者如果一些Actors改变了消息会怎么样?还记得Ruby不是一种不可变的语言吗?
- 与手动处理线程相比，Celluloid可以更快地构建复杂的并发系统。但是这是有[运行成本的](http://www.mikeperham.com/2015/10/14/optimizing-sidekiq/)(例如，减少5倍的速度和增多8倍的内存)。
- 不幸的是，Ruby实现并不擅长跨多个服务器使用分布式Actors。例如，使用0MQ的[DCell](https://github.com/loid/dcell)还没有准备好生产。

### 开源项目：

- [Reel](https://github.com/celluloid/reel/) — 基于事件的web服务器，它与基于`Celluloid`的应用程序一起工作。每个连接使用一个Actor。可以用于streaming或WebSockets。
- [Celluloid::IO](https://github.com/celluloid/celluloid-io/) — 将Actor和事件I/O循环结合在一起。与EventMachine不同，它允许通过创建多个Actor在每个进程中使用尽可能多的事件循环。 

## CSP(Communicating Sequential Processes)

Communicating Sequential Processes (CSP) 是一个非常类似于Actor模型的范例. 同样基于消息传递而不是共享内存。然而，CSP 和 Actors 有两个主要的区别：

- 进程在 CSP 中是匿名的，而 在Actors中是有标识的。所以，CSP 使用显式的通道来传递消息，而 Actors 是直接发送消息。
- CSP 中，发送者不能发送消息，直到接收者准备好接收前。Actors 可以异步发送消息（比如： [async calls](https://github.com/celluloid/celluloid/wiki/Basic-usage) in Celluloid）

CSP is implemented in such programming languages as Go with [goroutines and channels](https://blog.golang.org/share-memory-by-communicating), Clojure with the [core.async](http://clojure.com/blog/2013/06/28/clojure-core-async-channels.html) library and Crystal with [fibers and channels](https://crystal-lang.org/docs/guides/concurrency.html). For Ruby, there are a few gems which implement CSP. One of them is the `Channel` class implemented in [concurrent-ruby](https://github.com/ruby-concurrency/concurrent-ruby/blob/df482db36caf1b0c1d69a8ff97a2407469e1e315/doc/channel.md) library:

CSP 在 Go中使用了[goroutines and channels](https://blog.golang.org/share-memory-by-communicating)来实现，Clojure有[core.async](http://clojure.com/blog/2013/06/28/clojure-core-async-channels.html) 库，以及 Crystal 有 [fibers and channels](https://crystal-lang.org/docs/guides/concurrency.html)。对于 Ruby，有一些 gem 实现了 CSP。其中一个是在[concurrent-ruby](https://github.com/ruby-concurrency/conruby/blob/df482db36caf1d69a8ff97a2407469e1e315/doc/channel.md)库中实现的' Channel '类:

```ruby
# csp.rb
require 'concurrent-edge'
array = [1, 2, 3, 4, 5]
channel = Concurrent::Channel.new
Concurrent::Channel.go do
  puts "Go 1 thread: #{Thread.current.object_id}"
  channel.put(array[0..2].sum) # Enumerable#sum from Ruby 2.4
end
Concurrent::Channel.go do
  puts "Go 2 thread: #{Thread.current.object_id}"
  channel.put(array[2..4].sum)
end
puts "Main thread: #{Thread.current.object_id}"
puts channel.take + channel.take
$ gem install concurrent-ruby
$ gem install concurrent-ruby-edge
$ ruby csp.rb
Main thread: 70168382536020
Go 2 thread: 70168386894280
Go 1 thread: 70168386894880
18
```

因此，我们在2个不同的线程中运行2个操作(sum)，同步结果并在主线程中计算总数。所有操作都是通过通道完成的，没有任何显式的锁。

底层原理是，每个`Channel.go`都从线程池中一个独立的线程中运行，如果没有剩余的空闲线程，它会自动增加其大小。在这种情况下，在阻塞I/O操作期间使用这个模型非常有用，该操作释放了GIL(请参阅前面的文章以获得更多信息)。另一方面，`core.asynv`在Clojure中，异步使用有限数量的线程并试图“停放”它们，但是这种方法在I/O操作期间可能会出现问题，因为它可能会阻塞其他工作。

### 优点：

- CSP通道最多只能容纳一条消息，这使其更容易推理。而对于Actor模型，它更像是一个具有无限消息的邮箱。
- CSP允许您通过使用通道避免生产者和消费者之间的耦合;他们不需要了解彼此。
- 在CSP中，消息是按照它们被发送的顺序发送的。

> Clojure最终可能会支持分布式编程的actor模型，只在需要进行分发时才付出代价，但我认为它对于相同进程的编程非常麻烦。 [Rich Hickey](https://clojure.org/about/state#actors)

### 缺点：

- CSP is generally used on a single machine, it’s not that great as the Actor model for distributed programming.
- CSP通常使用在单台机器上，它没有分布式编程的Actor模型那么好。
- 在Ruby中，大多数实现都不使用M:N线程模型，因此每个“goroutine”实际上都使用一个Ruby线程，它等于一个OS线程。这意味着Ruby“goroutines”并不是轻量级的。
- 在Ruby中使用CSP不太流行。因此，目前还没有积极开发、稳定和经过实战检验的工具。

### 开源项目：

- [Agent](https://github.com/igrigorik/agent) — 另一个 Ruby 实现的 CSP。这个gem也在一个单独的Ruby线程中运行每个go-block。

## Software Transactional Memory

Actors和CSP是基于消息传递的并发模型，而Software Transactional Memory (STM)是使用共享内存的模型。它是基于锁的同步的另一种选择。与DB事务类似，这些是主要的概念:

1. 可以更改事务内的值，但是在**提交事务**之前，其他人无法看到这些更改。
2. 事务中发生的错误将**终止**它们并回滚所有更改。
3. 如果由于冲突的更改而无法提交事务，则会重试，直到成功。

[concurrent-ruby](https://github.com/ruby-concurrency/concurrent-ruby) gem实现了基于Clojure参考的[TVar](https://ruby-concurrency.github.io/concurrent-ruby/Concurrent/TVar.html).下面是一个例子，它实现了从一个银行账户到另一个银行账户的转账:

```ruby
# stm.rb
require 'concurrent'
account1 = Concurrent::TVar.new(100)
account2 = Concurrent::TVar.new(100)
Concurrent::atomically do
  account1.value -= 10
  account2.value += 10
end
puts "Account1: #{account1.value}, Account2: #{account2.value}"
$ ruby stm.rb
Account1: 90, Account2: 110
```

TVar是一个包含单个值的对象。它们与`atomically`一起在事务中实现数据更改。

### 优点：

- 使用 STM 比基于锁编程更简单。允许避免死锁，简化了对并发系统的推理，因为您不需要考虑竞争条件。
- 更容易适应，因为你不需要重组你的代码，如果使用 Actors和 CSP 是需要的。

### 缺点：

- 由于STM依赖于事务回滚，您应该能够在事务中的任何时间点撤消操作。在实践中，如果进行I/O操作(例如POST HTTP请求)，就很难保证。
- STM与Ruby MRI不能很好地匹配。由于有GIL，您只能使用一个CPU。同时，您也不能利用在线程中运行并发I/O操作的优势，因为很难撤消这些操作。

### 开源项目:

- TVar from [concurrent-ruby](https://github.com/ruby-concurrency/concurrent-ruby) — 实现STM并包含一些[基准测试](https://ruby-concurrency.github.io/conruby/concurrent/tvar.html)，分别在MRI、JRuby和Rubinius中使用STM和基于锁的实现来做了比较。

## Guilds

Guild是一个新的并发模型，由Ruby核心开发者Koichi Sasada为Ruby 3提出，他设计了当前的Ruby VM(虚拟机)、纤程和GC(垃圾收集器)。以下是创建Guilds的要点:

- 新模型应该与Ruby 2兼容，并允许更好的并发性。
- 强制使用与Elixir类似的不可变数据结构可能会慢得让人无法接受，因为Ruby使用了许多“写”操作。因此，最好复制与Racket ([Place](https://docs.racket-lang.org/reference/places.html))类似的共享可变对象，但复制必须快速才能成功。
- 如果有必要共享可变对象，那么应该有与Clojure(例如STM)类似的特殊数据结构。

这些想法导致了以下Guilds的主要概念:

- Guild是一个并发原语，它可以包含多个线程，也可以包含多个纤程。
- 只有Guild所有者可以访问它的可变对象，所以不需要使用锁。
- Guilds可以通过复制对象或将成员(“移动”对象)从一个Guild转移到另一个Guild来共享数据。
- 不可变对象可以通过引用访问任何工会，而不需要复制。例如numbers, symbols, `true`, `false`, deeply frozen objects.

因此，我们的钱从一个银行帐户转移到另一个的例子可能看起来像:

```ruby
bank = Guild.new do
  accounts = ...
  while acc1, acc2, amount, channel = Guild.default_channel.receive
    accounts[acc1].balance += amount
    accounts[acc2].balance -= amount
    channel.transfer(:finished)
  end
end
channel = Guild::Channel.new
bank.transfer([acc1, acc2, 10, channel])
puts channel.receive
# => :finished
```

所有关于账户余额的数据都存储在一个Guild(银行)中，因此，只有这个Guild负责可以通过通道请求的数据修改。

### 优点：

- Guilds之间没有可变的共享数据意味着不需要锁机制，因此没有死锁。Guilds之间的通信是为了安全而设计的。
- Guild鼓励使用不可变的数据结构，因为它们是跨多个Guild共享数据的最快和最简单的方法。现在开始冻结尽可能多的数据，例如，在文件的开头添加`# frozen_string_literal: true `。
- Guilds与Ruby 2完全兼容，这意味着您当前的代码将只在一个Guild内运行。您不需要使用不可变的数据结构或在代码中做任何更改。
- 与此同时，Guilds通过MRI实现了更好的并发性。它最终允许我们在一个Ruby进程中使用多个cpu。

### 缺点：

- It’s too early to make predictions about performance, but communicating and sharing mutable objects between Guilds will probably have a bigger overhead compared to threads.
- 现在预测性能还为时过早，但是与线程相比，在Guild之间通信和共享可变对象可能会有更大的开销。
- Guild是更复杂的并发原语，因为它们允许同时使用多个并发模型。例如:用于通过通道进行Guild间通信的CSP、具有特殊数据结构的STM(用于共享可变数据以获得更好的性能)、单个Guild内的多线程编程等等。
- 即使从资源使用的角度来看，在一个进程中运行多个Guild比运行多个进程开销要少，但是这个Guilds并不是轻量级的。它们将比Ruby线程更重，这意味着您无法仅使用Guild处理数万个WebSocket连接。

### 开源项目:

因为Ruby 3还没有发布，所以还没有例子。但是我看到了一个光明的未来，开发人员将开始构建对行业友好的工具，如web服务器、后台作业处理等。最有可能的是，所有这些工具都允许使用混合方法:运行多个进程、多个Guild和每个Guild中的多个线程。但现在，您可以阅读由佐田弘一(Koichi Sasada)撰写的原始[PDF演示文稿](http://www.atdot.net/~ko1/activities/2016_rubykaigi.pdf)。

# 结论

没有什么银弹。文章中描述的每个并发模型都有其优缺点。CSP模型在没有死锁的情况下在一台机器上运行得最好。Actor模型可以很容易地跨多台机器伸缩。STM允许编写更简单的并发代码。但是所有这些模型在Ruby中都不是一等公民，并且不能完全适应其他编程语言;主要是因为在Ruby中，它们都是用标准的并发原语(如线程和纤程)实现的。然而，有可能使用Ruby 3发布Guild，这是向更好的并发模型迈进了一大步!