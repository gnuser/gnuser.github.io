---
key: introduction-to-concurrency-models-with-ruby.-part-i
title: Introduction to Concurrency Models with Ruby. Part I
date: 2020-05-31 11:54:42 +0800
tags: ruby GIL concurrency parallel event-machine
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

在这第一篇文章中，我想描述进程、线程、GIL是什么、EventMachine和Ruby中的纤程之间的区别。什么时候使用哪个模型，哪个开源项目使用它们，优缺点是什么？

<!--more-->

## 多进程

运行多个进程实际上与并发(concurrency)无关，而是与并行(parallelism)有关。虽然并行和并发经常被混淆，但它们是不同的东西。我喜欢这个简单的类比:

- 并发性:让一个人只用一只手玩多个球。不管看起来如何，这个人一次只接/扔一个球。
- 并行性:让多人同时玩他们自己的一堆球。

### 顺序执行

想象一下，我们有一个数字范围，我们需要把它转换成一个数组，并为特定的元素找到一个索引:

```ruby
# sequential.rb
range = 0...10_000_000
number = 8_888_888
puts range.to_a.index(number)
$ time ruby sequential.rb                                                   
8888888
ruby test.rb  0.41s user 0.06s system 95% cpu 0.502 total
```

执行这段代码大约需要500毫秒，占用1个CPU。

### 并发执行

我们可以通过使用多个并行进程和分割范围来重写上面的代码。使用标准Ruby库中的fork方法，我们可以创建一个子进程并在块中执行代码。在父进程中，我们可以等待，直到所有子进程完成。

```ruby
# parallel.rb
range1 = 0...5_000_000
range2 = 5_000_000...10_000_000
number = 8_888_888
puts "Parent #{Process.pid}"
fork { puts "Child1 #{Process.pid}: #{range1.to_a.index(number)}" }
fork { puts "Child2 #{Process.pid}: #{range2.to_a.index(number)}" }
Process.wait
$ time ruby parallel.rb
Parent 32771
Child2 32867: 3888888
Child1 32865:
ruby parallel.rb  0.40s user 0.07s system 153% cpu 0.309 total
```

因为每个进程只使用范围的一半进行并行工作，所以上面的代码工作起来要快一些，并且要消耗1个以上的CPU。执行过程中的进程树如下：

```shell
# \ - 32771 ruby parallel.rb (parent process)
#  | - 32865 ruby parallel.rb (child process)
#  | - 32867 ruby parallel.rb (child process)
```

###优点：

- 进程不共享内存，因此你不能去修改另一个进程的数据。这可以让代码编写和调试更加容易。
- 由于有[Ruby MRI](https://en.wikipedia.org/wiki/Ruby_MRI)，多进程是唯一可以利用多核的方法，因为有一个GIL(全局解释器锁，下文有更多关于此的信息)。如果你在做比如说一些数学计算，这可能会很有用。
- 子进程可以避免不必要的内存泄漏。进程完成后，将释放所有资源。

### 缺点：

- 由于进程不共享内存，它们会使用大量内存——这意味着运行数百个进程可能是一个问题。注意，因为Ruby 2.0`fork`使用了OS [Copy-On-Write](https://en.wikipedia.org/wiki/Copy-on-write)，它允许进程共享内存，只要进程间没有不同的值。
- 进程创建和销毁很慢
- 多进程可能需要进程间通信。比如，[DRb](https://ruby-doc.org/stdlib-2.4.1/libdoc/drb/rdoc/DRb.html).
- 注意 [孤儿](https://en.wikipedia.org/wiki/Orphan_process) 进程(子进程的父进程结束或终止)或者[僵尸](https://en.wikipedia.org/wiki/Zombie_process)(子进程结束了但是仍然占据着进程表的空间)

### 开源项目：

- [Unicorn](https://bogomips.org/unicorn/)服务器——它加载应用程序，派生主进程以产生多个接受HTTP请求的工作进程。
- [Resque](https://github.com/resque/resque)用于后台处理——它运行一个worker，每个作业都顺序执行在一个fork的子进程中。

## 多线程

尽管Ruby从1.9版本开始就使用本地OS线程，但是在单个进程中，在任何给定时间都只能执行一个线程，即使您有多个cpu也是如此。这是因为MRI有GIL，而GIL也存在于Python等其他编程语言中。

### 为什么 GIL 存在

有几个原因，比如 :

- 避免C扩展中的竞争条件，无需担心线程安全。
- 易于实现，不需要使Ruby数据结构成为线程安全的。

早在2014年，Matz就开始考虑[逐渐移除GIL](https://twitter.com/yukihiro_matz/status/495219763883163648)。因为GIL实际上并不能保证我们的Ruby代码是线程安全的，也不允许我们使用更好的并发性。

### 竞态条件

下面是一个关于竞态条件的基本例子:

```ruby
# threads.rb
@executed = false
def ensure_executed
  unless @executed
    puts "executing!"
    @executed = true
  end
end
threads = 10.times.map { Thread.new { ensure_executed } }
threads.each(&:join)
$ ruby threads.rb
executing!
executing!
```

我们创建了10个线程来执行我们的方法，并为每个线程调用join，因此主线程将一直等待，直到所有其他线程都完成。代码打印`executing!`两次，因为我们的线程共享同一个@executed变量。我们的read(unless @executed)和set (@executed = true)操作不是原子性的，这意味着一旦我们读取了值，在我们设置新值之前，它可能在其他线程中被更改。

### GIL 和 阻塞I/O

但是拥有GIL(不允许同时执行多个线程)并不意味着线程就没有用处。当线程遇到阻塞I/O操作会释放 GIL，如HTTP请求，DB查询，读写磁盘，甚至睡眠:

```ruby
# sleep.rb
threads = 10.times.map do |i|
  Thread.new { sleep 1 }
end
threads.each(&:join)
$ time ruby sleep.rb                                                    
ruby sleep.rb  0.08s user 0.03s system 9% cpu 1.130 total
```

如你所见，所有10个线程都休眠了1秒钟，并且几乎同时完成。当一个线程进入睡眠状态时，它将执行传递给另一个线程，而不阻塞GIL。

### 优点:

- 使用的内存比进程少;可以运行数千个线程。它们创造和销毁的速度也很快。
- 如果有慢的阻塞 IO 操作，线程会很有用
- 可以访问其他线程的内存空间

### 缺点:

- 需要非常小心的同步以避免竞争条件，通常通过使用锁原语，这有时可能导致死锁。所有这些都使得编写、测试和调试线程安全代码变得非常困难。
- 对于线程，必须确保不仅代码是线程安全的，而且使用的任何依赖项也是线程安全的。
- 派生的线程越多，它们通过切换上下文花费的时间和资源就越多，而完成实际工作的时间就越少。

### 开源项目：

- [Puma](https://github.com/puma/puma)  - 允许在每个进程中使用多个线程(集群模式)。与Unicorn类似，它预加载应用程序并派生主进程，其中每个子进程都有自己的线程池。在大多数情况下，线程工作得很好，因为每个HTTP请求都可以在一个单独的线程中处理，而且我们不会在请求之间共享很多资源。

- [Sidekiq](https://github.com/mperham/sidekiq) - 用于后台处理 - 默认情况下运行一个带有25个线程的进程。每个线程一次处理一个任务。

  

### EventMachine

EventMachine(又名EM)是一个用c++和Ruby编写的gem。它使用 [Reactor pattern](https://en.wikipedia.org/wiki/Reactor_pattern)提供事件驱动的I/O，并且基本上可以使你的Ruby代码看起来像Node.js:)。EM在运行事件循环期间使用Linux select()检查文件描述符上的新输入。

使用EventMachine的一个常见原因是，如果有很多I/O操作，并且不希望手动处理线程。从资源使用的角度来看，手动处理线程可能很困难，或者成本通常太高。而使用 EM就可以在默认情况下用一个线程处理多个HTTP请求。

```ruby
# em.rb
EM.run do
  EM.add_timer(1) do
    puts 'sleeping...'
    EM.system('sleep 1') { puts "woke up!" }
    puts 'continuing...'
  end
  EM.add_timer(3) { EM.stop }
end
$ ruby em.rb
sleeping...
continuing...
woke up!
```

![render1590909300420](/../../../../../../../media/2020-05-31-introduction-to-concurrency-models-with-ruby.-part-i/render1590909300420.gif)

上面的示例显示了如何通过执行EM.system (I/O操作)并传递一个作为回调的块来运行异步代码，回调将在系统命令完成后执行。

### 优点：

- 可以使用单线程提升 web server 和 proxies 的性能
- 避免复杂的多线程编程

### 缺点：

- 每个 I/O 操作都要支持异步的 EM。意味着你必须使用特定版本的系统，数据库适配器，HTTP 客户端等等。这会导致补丁版本，缺乏支持和选项限制
- 在主线程中，每个事件循环所做的工作应该尽量少。同样，也可以使用[Defer](http://www.rubydoc.info/github/eventmachine/eventmachine/EventMachine.defer)，它在与线程池分开的线程中执行代码，但是，这可能会导致前面讨论的多线程问题。
- 由于错误处理和回调，很难编写复杂的系统。回调地狱在Ruby中也是可能的，但是可以通过使用纤程来阻止它，如下所示。
- EventMachine本身是一个巨大的依赖:Ruby中有17K LOC(代码行数)，c++中有10K LOC。

### 开源项目：

- [Goliath](https://github.com/postrank-labs/goliath/) — 单线程异步服务器
- [AMQP](https://github.com/ruby-amqp/amqp) — RabbitMQ 客户端. 然而，这个 gem 的作者建议使用不基于 EM 的版本[Bunny](http://rubybunny.info/). 请注意，将工具迁移到没有em的实现是一种普遍趋势。例如,(ActionCable) (https://github.com/rails/rails/tree/master/actioncable)的作者决定使用低级(nio4r) (https://github.com/socketry/nio4r)，(sinatra-synchrony) (https://github.com/kyledrake/sinatra-synchrony)的作者使用[Celluloid](https://github.com/celluloid/celluloid)重写了,等等。

## Fibers(纤程)

[Fibers](https://ruby-doc.org/co-2.4.1/fiber.html)是Ruby标准库中的轻量级原语，可以手动暂停、恢复和调度。如果你熟悉JavaScript，那么它们与ES6生成器非常相似，我们还写了一篇关于[Generators and Redux-Saga](https://engineering.universe.com/what-redux -saga-c1252fc2f4d1)的文章 。在一个线程中可以运行数万个纤程。

通常，在EventMachine中使用纤程可以避免回调并使代码看起来同步。所以，下面的代码:

```ruby
EventMachine.run do
  page = EM::HttpRequest.new('https://google.ca/').get       
  page.errback { puts "Google is down" }
  page.callback {
    url = 'https://google.ca/search?q=universe.com'
    about = EM::HttpRequest.new(url).get
    about.errback  { ... }
    about.callback { ... }     
  }
end
```

可以重写成：

```ruby
EventMachine.run do
  Fiber.new {
    page = http_get('http://www.google.com/')     
    if page.response_header.status == 200
      about = http_get('https://google.ca/search?q=universe.com') 
      # ... 
    else 
      puts "Google is down"
    end  
  }.resume 
end
def http_get(url)
  current_fiber = Fiber.current
  http = EM::HttpRequest.new(url).get    
  http.callback { current_fiber.resume(http) }   
  http.errback  { current_fiber.resume(http) }    
  Fiber.yield
end
```

因此，基本上，Fiber#yield 返回到恢复纤程的上下文，并把返回值传递给Fiber#resume。

### 优点

- Fibers allow you to simplify asynchronous code by replacing nested callbacks.
- 纤程允许您通过替换嵌套回调来简化异步代码。

### 缺点:

- 并不能真正解决并发问题。
- 它们很少直接在应用程序级代码中使用

### 开源项目：

- [em-synchrony](https://github.com/igrigorik/em-synchrony) — 一个库，作者Ilya Grigorik, 是一名在 Google 的性能工程师, 它为不同的客户端(如MySQL2、Mongo、Memcached等)集成了EventMachine和纤程。

## 结论

没有什么银弹，所以要根据自己的需要选择并发模型。例如：

- 需要运行CPU和内存密集型代码，并有足够的资源使用多进程。
- 必须执行多个I/O操作，如HTTP请求-使用多线程。
- 需要扩大到最大的吞吐量-使用EventMachine。

在本系列的第二部分中，我们将研究actor (Erlang, Scala)、通信顺序进程(Go, Crystal)、软件事务性内存(Clojure)等并发模型，当然还有Guilds——一种可能在Ruby 3中实现的新并发模型。

