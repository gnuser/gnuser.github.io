---
key: factory_bot的transient用法
title: FactoryBot的transient用法
date: 2020-07-17 15:02:33 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

FactoryBot的transient用法

<!--more-->

### 用来控制别的属性

```ruby
factory :user do
  transient do
    rockstar { true }
  end

  name { "John Doe#{" - Rockstar" if rockstar}" }
end

create(:user).name
#=> "John Doe - ROCKSTAR"

create(:user, rockstar: false).name
#=> "John Doe"
```

### 当使用`attributes_for`方法

放在`transient`里的属性会被忽略

### 回调时访问

```ruby
factory :invoice do
  trait :with_amount do
    transient do
      amount { 1 }
    end

    after(:create) do |invoice, evaluator|
      create :line_item, invoice: invoice, amount: evaluator.amount
    end
  end
end

create :invoice, :with_amount, amount: 2
```

### 关联数据时

不支持这样的`transient`关联，所以即使放在`transient`里，关联数据也算是正常的数据，`non-transient`的数据

```ruby
factory :post

factory :user do
  transient do
    post { build(:post) }
  end
end
```





* [https://www.rubydoc.info/gems/factory_bot/file/GETTING_STARTED.md#transient-attributes](https://www.rubydoc.info/gems/factory_bot/file/GETTING_STARTED.md#transient-attributes)