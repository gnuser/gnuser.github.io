---
key: how-to-love-ruby
title: how-to-love-ruby
date: 2020-09-01 14:55:11 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

快速学习Ruby ，从此爱上这门语言

<!--more-->

> I like to program in Ruby because Ruby is a *succinct* and is a highly *practical* language that is *widely used* by many programmers in the industry.

用 Ruby 编写代码可以非常简洁。

## Ruby 生态

https://www.jetbrains.com/lp/devecosystem-2019/ruby/

ruby 版本管理器: 

- rbenv
- rvm

ruby 版本: 

- 2.6
- 2.5 (most)

rails 版本：

- edge
- 5.2 (most)
- 5.1
- 5.0
- 4.2

web framework 框架：

- Rails (most)
- Rack
- Sinatra

server:

- puma (most)
- passenger
- unicorn
- thin

IDE:

- RubyMine (most)
- VS Code

unit-test framework:

- RSpec (most)
- MiniTest

## 20 分钟教程

https://www.ruby-lang.org/en/documentation/quickstart/

### 安装

先安装好 ruby，如果没安装过，参考[https://rvm.io/](https://rvm.io/)

### irb

终端输入 `irb`，就可以开始编写简单的 ruby 程序了

```
 :001 > 'hello'
 => "hello"
 :003 > 'world'
 => "world"
 :004 > 1+ 2
 => 3
 :005 > 12312312312313131231313123 + 12423431342134123424241243123
 => 12435743654446436555472556246
```

### 函数

函数定义

```ruby
def hi
  puts 'hello world'
end
```

调用函数时，可以不写括号，即使有参数也可以不写括号

```ruby
irb(main):013:0> hi
Hello World!
=> nil
irb(main):014:0> hi()
Hello World!
=> nil
```

```ruby
 :010 > def hi(name, age)
 :011?>   puts "hello, #{name}, you are #{age}"
 :012?>   end
 => :hi
 :013 > hi 'cj', 18
hello, cj, you are 18
 => nill
```

字符串内的`#{name}`可以写更复杂的语句，比如`#{name.capitalize}`

函数的默认参数使用

```ruby
irb(main):019:0> def hi(name = "World")
irb(main):020:1> puts "Hello #{name.capitalize}!"
irb(main):021:1> end
=> :hi
irb(main):022:0> hi "chris"
Hello Chris!
=> nil
irb(main):023:0> hi
Hello World!
=> nil
```

使用符号类型参数时，调用时，需要连同符号一起传入

```ruby
 :024 > def hi(name: 'world')
 :025?>   puts name
 :026?> end
 => :hi
  :029 > hi name: 'cj'
cj
 => nil
```

### 类

实例变量

```ruby
irb(main):024:0> class Greeter
irb(main):025:1>   def initialize(name = "World")
irb(main):026:2>     @name = name
irb(main):027:2>   end
irb(main):028:1>   def say_hi
irb(main):029:2>     puts "Hi #{@name}!"
irb(main):030:2>   end
irb(main):031:1>   def say_bye
irb(main):032:2>     puts "Bye #{@name}, come back soon."
irb(main):033:2>   end
irb(main):034:1> end
=> :say_bye

irb(main):035:0> greeter = Greeter.new("Pat")
=> #<Greeter:0x16cac @name="Pat">
irb(main):036:0> greeter.say_hi
Hi Pat!
=> nil
irb(main):037:0> greeter.say_bye
Bye Pat, come back soon.
=> nil
```

类的方法列表

```ruby
irb(main):039:0> Greeter.instance_methods
=> [:say_hi, :say_bye, :instance_of?, :public_send,
    :instance_variable_get, :instance_variable_set,
    :instance_variable_defined?, :remove_instance_variable,
    :private_methods, :kind_of?, :instance_variables, :tap,
    :is_a?, :extend, :define_singleton_method, :to_enum,
    :enum_for, :<=>, :===, :=~, :!~, :eql?, :respond_to?,
    :freeze, :inspect, :display, :send, :object_id, :to_s,
    :method, :public_method, :singleton_method, :nil?, :hash,
    :class, :singleton_class, :clone, :dup, :itself, :taint,
    :tainted?, :untaint, :untrust, :trust, :untrusted?, :methods,
    :protected_methods, :frozen?, :public_methods, :singleton_methods,
    :!, :==, :!=, :__send__, :equal?, :instance_eval, :instance_exec, :__id__]
```

上面输出了包括父类的所有方法，如果我们只想打印当前类的方法，传入参数`false`,表示我们不需要父类方法

```ruby
irb(main):040:0> Greeter.instance_methods(false)
=> [:say_hi, :say_bye]
```

可以使用`respond_to?`来查看是否支持该方法

```ruby
irb(main):041:0> greeter.respond_to?("name")
=> false
irb(main):042:0> greeter.respond_to?("say_hi")
=> true
irb(main):043:0> greeter.respond_to?("to_s")
=> true
```

我们可以方便的添加`get`, `set`方法，使用`attr_accessor`

```ruby
class Greeter
  attr_accessor :name

  def initialize(name = 'world')
    @name = name
  end

  def say_hi
    puts @name
  end
end

greeter = Greeter.new
greeter.name = 'cj'
pp greeter.name
```

动态修改成员变量

```ruby
#!/usr/bin/env ruby

class MegaGreeter
  attr_accessor :names

  # Create the object
  def initialize(names = "World")
    @names = names
  end

  # Say hi to everybody
  def say_hi
    if @names.nil?
      puts "..."
    elsif @names.respond_to?("each")
      # @names is a list of some kind, iterate!
      @names.each do |name|
        puts "Hello #{name}!"
      end
    else
      puts "Hello #{@names}!"
    end
  end

  # Say bye to everybody
  def say_bye
    if @names.nil?
      puts "..."
    elsif @names.respond_to?("join")
      # Join the list elements with commas
      puts "Goodbye #{@names.join(", ")}.  Come back soon!"
    else
      puts "Goodbye #{@names}.  Come back soon!"
    end
  end
end


if __FILE__ == $0
  mg = MegaGreeter.new
  mg.say_hi
  mg.say_bye

  # Change name to be "Zeke"
  mg.names = "Zeke"
  mg.say_hi
  mg.say_bye

  # Change the name to an array of names
  mg.names = ["Albert", "Brenda", "Charles",
              "Dave", "Engelbert"]
  mg.say_hi
  mg.say_bye

  # Change to nil
  mg.names = nil
  mg.say_hi
  mg.say_bye
end
```

```ruby
Hello World!
Goodbye World.  Come back soon!
Hello Zeke!
Goodbye Zeke.  Come back soon!
Hello Albert!
Hello Brenda!
Hello Charles!
Hello Dave!
Hello Engelbert!
Goodbye Albert, Brenda, Charles, Dave, Engelbert.  Come
back soon!
...
...
```

### 字符串和符号(symbol)

双引号和单引号的区别？

字符串转换大小写?

判断字符串是否为空？

什么时候用 symbol? 

```ruby
"string_to_symbol".to_sym # => :string_to_symbol
"string to symbol".to_sym # => :"string to symbol"
symbol_to_string.to_s     # => "symbol_to_string"
```

### ruby 和 rails中的各种奇怪符号

- `$`: global variables are available everywhere in a program.
- `@`: instance variables
- `@@`: class variables
- `&`:
- `:` 
- `?`
- `*`
- `!`: make a really change

### ruby 中的容器遍历

这里的 do 和 end 中间称为`block`，类似于`anonymous function` or `lambda`，两根竖线

```ruby
@names.each do |name|
  puts "Hello #{name}"
end
```

相对于 C 语言，ruby is more elegant~

```C
for (i=0; i<number_of_elements; i++)
{
  do_something_with(element[i]);
}
```

将数组中的字符串使用逗号连接输出

```ruby
[1,2,3].join(",")
```

### enumrable

https://docs.ruby-lang.org/en/2.6.0/Enumerable.html



ruby[官方文档](https://docs.ruby-lang.org/en/2.6.0)

