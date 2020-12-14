---
key: whats-the-difference-between-load-and-require
title: whats-the-difference-between-load-and-require
date: 2020-12-14 20:28:33 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

load和require的区别，gems和bundler以及rails autoload机制

<!--more-->

## load vs require



| load                                         | require                       |
| -------------------------------------------- | ----------------------------- |
| each time will execute the code              | only execute once             |
| need to pass extension '.rb'                 | no need to pass the extension |
| will check the file in the current directory | only search the $LOAD_PATH    |

## gems

1. get the gem version

```shell
gem list eth
*** LOCAL GEMS ***

eth (0.4.12)
method_source (0.9.2)
```

2. install the specific version

```shell
gem install domain_name -v 0.5.20160826
```

3. find the gem location

```shell
gem which eth
/Users/gnuser/.rvm/gems/ruby-2.6.5/gems/eth-0.4.12/lib/eth.rb
```

## bundler

`bundle exec rspec` will use the gem version specified in `Gemfile.lock` 

## Rails load all gems flow

1. `config/boot.rb`

```ruby
ENV['BUNDLE_GEMFILE'] ||= File.expand_path('../Gemfile', __dir__)
require 'bundler/setup' # Set up gems listed in the Gemfile
```

`setup` will read `Gemfile.lock` to call `gem` method for each gem with the correct version, if the gem version mismatched will raise an exception.

2. `application.rb`

```ruby
# Require the gems listed in Gemfile, including any gems
# you've limited to :test, :development, or :production.
Bundler.require(:default, Rails.env)
```

call `kernel.require` for each gem.

### require: false in Gemfile meaning

by default, `Bundler.require` will require every gem from Gemfile. if we use `require: false`, you can explictly `require` the gem when you want to use it.



 

