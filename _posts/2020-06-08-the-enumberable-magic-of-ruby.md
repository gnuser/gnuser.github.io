---
key: the-enumberable-magic-of-ruby
title: The Enumberable magic of Ruby
date: 2020-06-08 23:39:26 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

今天看到几行代码，发现 Ruby 真的是处处有惊喜，那简约的风格真的能让代码成为艺术品！

<!--more-->

之前有行代码类似这样:

```ruby
EVENT_TYPE_LIST = %w(Accounts::PaymentDepositCredited Accounts::PaymentRefundRequested Accounts::PaymentRefundOngoingChecked Accounts::PaymentRefundTransferred Accounts::PayoutCreated Accounts::PayoutProcessed Accounts::PayoutRejected Accounts::PayoutFailed Accounts::ConversionConfirmed)
```

分行展示吧，对齐也着实难看，还隐隐感觉分了行会有问题，就放一行了，然后看到别人写的代码：

```ruby
EVENT_TYPE_LIST = %w(PaymentDepositCredited PaymentRefundRequested PaymentRefundOngoingChecked PaymentRefundTransferred PayoutCreated PayoutProcessed PayoutRejected PayoutFailed ConversionConfirmed).collect { |e| "Accounts::#{e}" }
```

嗯？连字符串的重复都不放过？我服了！

然后没完，接下来还有几行代码：

```ruby
EventRecord.where(event_type: t).each.with_object({}) {|e, h| h[e.id] ||= []; h[e.id] << e}.collect {|aid, events| Accounts::TransactionProjector.project(events.first.aggregate_id, events: events)}.count
```

shit,这一长串操作我反复看了快 半个小时，才知道他大概意思，当然可以写得更清晰好懂一点，但这么一行运行下来，其实没有太多拖泥带水的语句，反而还是体现了简洁的风格，所以瞬间感觉到了差距了，赶紧恶补一下 Ruby 语法。

## Enumerable

这里其实主要用到了 Ruby的`Enumerable module`,也就是迭代器，一般我们会用到三个迭代器: `each`, `collect`, `map`

比较常用的是 each,一般也就是遍历数组或者 hash，并不会改变原数据，而 collect 和 map 都会改变原来的数据并且直接返回新数据,而一般我们用`map`就可以了，因为`collect`和`map`是一样的

## with_object

还有这个 `with_object`,以前也比较少用，下面两条语句的作用也一样,`with_object`的参数相当于初始化了一个 object，然后传给后面的 `block`迭代使用

```ruby
p %w(foo bar).map.with_object({}) { |str, hsh| hsh[str] = str.upcase }
p %w(foo bar).each.with_object({}) { |str, hsh| hsh[str] = str.upcase }
```

