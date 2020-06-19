---
key: ruby-times
title: Ruby times
date: 2020-06-19 22:46:01 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

当你想写一个循环的时候，ruby 提供了最简单的方式 times

<!--more-->

目标只是简单的想循环 N 次，也可能因为一些条件break 跳出

刚开始的时候用的 util

```ruby
cd = 0
index = 0
until cd > 10
  p "before #{cd}, #{index}"
  next if index == 2
  index += 1
  p "after #{cd}, #{index}"
  cd += 1
end
```

这里如果条件写的不好,就死循环了。

所以如果希望最大的循环次数不要超过 N，那使用 times 会更清晰一些

```ruby
index = 0
10.times do
  p index
  next if index == 2
  index += 1
end
```

