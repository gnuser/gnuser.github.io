---
key: rails-convert-hash-key-from-camelcase-to-snakecae-and-vice-versa
title: rails-convert-hash-key-from-camelcase-to-snakecae-and-vice-versa
date: 2020-12-23 17:50:36 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

方便的转换hash key，驼峰和蛇形命名
<!--more-->

```ruby
obj.deep_transform_keys! { |key| key.camelcase }
```

```ruby
obj.deep_transform_keys! { |key| key.snake_case }
```
