---
key: factory-bot使用
title: FactoryBot使用
date: 2020-07-19 21:04:58 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

测试驱动开发中，需要构造一些数据，`FactoryBot`是常用的数据构造工具，语法简洁，支持多种数据构造策略，本文介绍`FactoryBot`在 rails 框架配合`RSpec`使用。

<!--more-->

### 安装

添加`gem`到`Gemfile`

```ruby
group :development, :test do
  gem "factory_bot_rails"
end
```

添加代码到`spec/support/factory_bot.rb`

```ruby
# spec/support/factory_bot.rb
RSpec.configure do |config|
  config.include FactoryBot::Syntax::Methods
end
```

激活自动加载`support`目录, 修改`spec/rails_helper.rb`

```ruby
Dir[Rails.root.join('spec/support/**/*.rb')].each { |f| require f }
```

### 使用案例

加入我们有一个`Article`的`model`：

- 有一个字段`status`，类型为`Integer`，表示文章的状态，并且可以取的值为`unpublished`和`published`,
- 还有个字段`published_at`,类型为`DateTime`，表示文章的发布时间

```ruby
# app/model/article.rb
class Article < ApplicationRecord
  enum status: [:unpublished, :published]
end
```

所以我们可以定义以下测试数据：

```ruby
# spec/factories/articles.rb
FactoryBot.define do
  factory :article do
    trait :published do
      status :published
    end

    trait :unpublished do
      status :unpublished
    end

    trait :in_the_future do
      published_at { 2.days.from_now }
    end

    trait :in_the_past do
      published_at { 2.days.ago }
    end
  end
end
```

现在我们可以在测试用例中构建数据：

```ruby
# build creates an Article object without saving
build :article, :unpublished

# build_stubbed creates an Article object and acts as an already saved Article
build_stubbed :article, :published

# create creates an Article object and saves it to the database
create :article, :published, :in_the_future
create :article, :published, :in_the_past

# create_list creates a collection of objects for a given factory
# you can also use build_list and build_stubbed_list
create_list :article, 2
```

可以看到构建数据的策略有:

- build: 只创建对象，但不入库，会触发`model`的`validations`
- build_stubbed：创建对象，不入库，也不触发任何 `validations`
- create: 创建对象，同时入库，会触发`model`和数据库的`validations`
- create_list: `create`的批量创建方式
- build_stubbed_list：`build_stubbed`的批量创建方式
- attributes_for: 获取构造对象的属性列表

### 最佳实践

- 使用Factory linting
- 填充必要的字段即可
- 尽量使用 build 和 build_stubbed而不是 create
- 不要依赖默认的数据定义来进行测试
- 使用固定时间而不是相对时间

#### Factory linting

安装`database_cleaner`

```ruby
gem :database_cleaner, group: :test
```

添加代码到`lib/tasks/factory_bot.rake`

```ruby
namespace :factory_bot do
  desc "Verify that all FactoryBot factories are valid"
  task lint: :environment do
    if Rails.env.test?
      DatabaseCleaner.cleaning do
        FactoryBot.lint
      end
    else
      system("bundle exec rake factory_bot:lint RAILS_ENV='test'")
      fail if $?.exitstatus.nonzero?
    end
  end
end
```

执行检查

```ruby
bundle exec rake factory_bot:lint
```

#### 填充必要的字段即可

比如一些状态字段，需要某些条件才能触发改变，就不应该放在数据工厂定义里

```ruby
# spec/factories/articles.rb
FactoryBot.define do
  factory :article do
    title "The amazing article title"

    trait :with_publish_date do
      published_at { DateTime.now }
    end
  end
end
```

这里我们可以使用任意选择是否需要构造带`published_at`的对象

```ruby
let(:article_with_publish_date) { build :article, :with_publish_date }
let(:article_without_publish_date) { build :article }
```

#### 不要依赖默认的数据定义来进行测试

```ruby
require 'rails_helper'

RSpec.describe Article do
  describe ".published_in_the_past" do
    before do
      create :article, title: 'unpublished article'
      create :article, :published, :in_the_past, title: 'published in the past'
      create :article, :published, :in_the_future, title: 'published in the future'
    end

    subject(:article_titles) { Article.published_in_the_past.map(&:title) }

    it { expect(article_titles).to include 'published in the past' }
    it { expect(article_titles).not_to include 'unpublished article' }
    it { expect(article_titles).not_to include 'published in the future' }
  end
end
```

#### 尽量使用 build 和 build_stubbed而不是 create

因为 create 会入库，如果滥用，只会让测试用例越来越慢，所以除非必要，尽量用 build 和 build_stubbed

因为 build 会把关联数据入库，如果我们不想入库，就使用`build_stubbed`来构造数据

```ruby
# spec/factories/articles.rb
FactoryBot.define do
  factory :article do
    name 'The amazing article'
    author
  end
end

# spec/models/article_spec.rb
require 'rails_helper'

RSpec.describe Article do
  describe ".recent" do
    let(:latest)   { build_stubbed :article, :published, title: :latest  }
    let(:promoted) { build_stubbed :article, :published, title: :promoted }
  end
end
```

#### 使用固定时间而不是相对时间

`2.seconds.ago`,`5.minutes.ago`可能会导致一些问题，比如统计类逻辑，所以最好用固定时间，比如

```ruby
create :article, published_at: "2015-04-04T17:30:05+0700"
```

或者使用工具来冻住时间，比如添加`ActiveSupport::Testing::TimeHelpers`到`RSpec`的配置中

```ruby
# spec/rails_helper.rb
RSpec.configure do |config|
  config.include ActiveSupport::Testing::TimeHelpers
end
```

可以这样使用:

```ruby
before do
  travel_to Time.current
end

after do
  travel_back
end
```

更多的最佳实践使用方法可以参考[https://github.com/thoughtbot/factory_bot/blob/master/GETTING_STARTED.md](https://github.com/thoughtbot/factory_bot/blob/master/GETTING_STARTED.md)



参考自 [https://semaphoreci.com/community/tutorials/working-effectively-with-data-factories-using-factorygirl](https://semaphoreci.com/community/tutorials/working-effectively-with-data-factories-using-factorygirl)