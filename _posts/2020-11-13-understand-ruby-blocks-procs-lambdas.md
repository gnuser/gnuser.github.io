---
key: understand-ruby-blocks,-procs-&-lambdas
title: Understand ruby blocks, procs & lambdas
date: 2020-11-13 23:55:33 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

什么是 blocks，procs 和 lambdas，如何使用？原理是什么？
<!--more-->

## 什么是 blocks

Ruby的blocks类似于匿名函数,并且能作为参数传给其他函数.

Blocks有两种写法：
- do end
- {}

这两种都可以传递多个参数.举个例子
```ruby
# {} 写法
[1, 2, 3].each { |i| p i }

# do end 写法
[1, 2, 3].each do |i|
  p i
end
```

## 什么是 yield

yield就是为了调用 blocks 的

举个例子
```ruby
def show
  yield
end

show { p 'hello' }
```

yield 还可以多次执行
```ruby
def show
  yield
  yield
  yield
end

show { p 'hello' }
```

yield 也可以传参，并且支持多个参数的传递，参数是传递到 blocks 中去的
```ruby
def get_balance
  yield 1
  yield 2
  yield 3
end

get_balance { |i| p i * 10000 }
```

## blocks 可以显式的传递

上面的 yield 方式算是隐式的传递

```ruby
def explict_block(&block)
  block.call # same as yield
end

explict_block { p 'hello' }
```

## 如何判断是否传入了 block

如果不传block,执行 yield 会报错: `no block given (yield)`

可以通过`block_given?`方法判断

```ruby
def implict_block
  return 'no block passed' unless block_given?
  yield
end

implict_block
```

## 什么是 lambda
lambda 可以定义好一个 block 以及其需要的参数.

语法是这样的:
```ruby
say_something = -> { puts "This is a lambda" }
```

> 可以使用`lambda`替换`->`

调用可以用`call`
```ruby
say_something = -> { puts "This is a lambda" }
say_something.call
```

也可以传递参数给 lambda
```ruby
sum = ->(a, b) { a + b }
p sum.call(1, 2)
```

## 什么是 procs

lambda 和 procs 类似，只是 procs 的定义方式不同,lambda是一个特殊的 `Proc`对象,
`Proc`传递参数的方式也不一样

```ruby
my_proc = Proc.new { |x| p x }
```

`Proc`的 `return`和 lambda 也不一样，lambda的`return`是退出自己，不影响调用方
而`Proc`的`return`是从当前方法中退出,跟直接在函数执行`return`一样

```ruby
def call_proc
  puts "Before proc"
  my_proc = Proc.new { return 2 }
  my_proc.call
  puts "After proc"
end
p call_proc
# Prints "Before proc" but not "After proc"
```

## Closures

Ruby 的`Proc`,`lambda`都有一个特殊的概念

这个概念叫闭包，可以访问到当前上下文的变量和方法，但是在闭包中即使修改了变量值，对上下文也不会影响

```ruby
def call_proc(my_proc)
  count = 500
  my_proc.call
end
count   = 1
my_proc = Proc.new { puts count }
call_proc(my_proc)
p count  # What does this print?
```

这里虽然`call_proc`修改了 `count`，也并不会修改到外面的`count`值











