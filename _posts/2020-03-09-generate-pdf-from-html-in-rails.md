---
key: generate-pdf-from-html-in-rails
title: 在rails中从html生成pdf
date: 2020-03-09 21:41:38 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

因为工作需要,要生成一个发票的pdf,这个需求还挺有意思的,要注意的地方比如分页的处理,还有不要用复杂的flex布局.

<!--more-->

目前的解决方案有:

- PDFKit
- Wicked PDF

这两个都是基于一个跨平台的免费开源工具[wkhtmltopdf](https://github.com/wkhtmltopdf/wkhtmltopdf).这个工具能直接从html页面生成pdf.

## 安装wkhtmltopdf

[下载对应平台的包](https://wkhtmltopdf.org/downloads.html)

安装好后,可以直接在命令行测试一下

```shell
 ✗ wkhtmltopdf https://gnuser.github.io/ gnuser.pdf
Loading pages (1/6)
Counting pages (2/6)
Resolving links (4/6)
Loading headers and footers (5/6)
Printing pages (6/6)
Done
```

这条指令会在当前目录根据博客首页生成gnuser.pdf文件.

![image-20200309222930139](/../../../../../../../media/2020-03-09-generate-pdf-from-html-in-rails/image-20200309222930139.png)

还带目录,感觉很棒吧!

## 新建一个rails项目

```shell
rails new rails-generate-pdf
```



## 安装wicked_pdf

添加gem到Gemfile

```ruby
gem 'wicked_pdf'
gem 'wkhtmltopdf-binary' # 如果上面一步已经安装过,这里可以注释掉.
```

执行

```shell
# command line
bundle install
```

## 生成初始化文件

```shell
# command line
✗ rails g wicked_pdf
Running via Spring preloader in process 72924
      create  config/initializers/wicked_pdf.rb
```

## 创建data models

- 生成Invoice模型

```shell
rails generate model Invoice from_full_name from_address from_email from_phone to_full_name to_address to_email to_phone status discount:decimal vat:decimal
```

修改`app/models/invoice.rb`文件

```ruby
# file: rails-generate-pdf/app/models/invoice.rb
class Invoice < ApplicationRecord
    has_many :invoice_items, dependent: :destroy

    STATUS_CLASS = {
        :draft => "badge badge-secondary",
        :sent => "badge badge-primary",
        :paid => "badge badge-success"
    }

    def subtotal
        self.invoice_items.map { |item| item.qty * item.price }.sum
    end

    def discount_calculated
        subtotal * (self.discount / 100.0)
    end

    def vat_calculated
        (subtotal - discount_calculated) * (self.vat / 100.0)
    end

    def total
        subtotal - discount_calculated + vat_calculated
    end

    def status_class
        STATUS_CLASS[self.status.to_sym]
    end

end
```

- 生成InvoiceItem模型

```shell
rails generate model InvoiceItem name description price:decimal qty:integer invoice:references
```

## 创建数据库

```shell
rails db:create
ralis db:migrate
```

## 生成测试数据

修改[db/seeds.rb](https://github.com/PDFTron/rails-generate-pdf/blob/master/db/seeds.rb)

```shell
rails db:seed
```

创建controller

```shell
rails generate controller Invoices index show
```

修改`app/controllers/invoices_controller.rb`

```ruby
# file: rails-generate-pdf/app/controllers/invoices_controller.rb
class InvoicesController < ApplicationController
    def index
        @invoices = scope
    end

    def show
        @invoice = scope.find(params[:id])

        respond_to do |format|
            format.html
            format.pdf do
                render pdf: "Invoice No. #{@invoice.id}",
                page_size: 'A4',
                template: "invoices/show.html.erb",
                layout: "pdf.html",
                orientation: "Landscape",
                lowquality: true,
                zoom: 1,
                dpi: 75
            end
        end
    end

    private
        def scope
            ::Invoice.all.includes(:invoice_items)
        end
end
```

这里我们在show方法实现了两种渲染方式(html和pdf), 只要访问`.pdf`后缀,我们就可以渲染pdf了.

可以试试这两个链接:

- http://rails-generate-pdf.herokuapp.com/invoices/1
- http://rails-generate-pdf.herokuapp.com/invoices/1.pdf

## 配置参数

- layout布局`pdf.html`: `app/views/layouts/pdf.html.erb`
- template文件`invoices/show.html.erb`
- page-size: A4纸
- orientation: Landscape或者Portrait(默认)
- lowquality: 降低pdf质量,压缩文件大小
- dpi: 修改dpi系数
- zoom: 缩放比例

## 添加路由

修改`app/config/routes.rb`

```ruby
# file: rails-generate-pdf/app/config/routes.rb
Rails.application.routes.draw do
    root to: 'invoices#index'

    resources :invoices, only: [:index, :show]
end
```

## 修改layout文件

```ruby
# file: rails-generate-pdf/views/layous/pdf.html.erb
<!DOCTYPE html>
<html>
<head>
<title>PDFs - Ruby on Rails</title>
    <%= wicked_pdf_stylesheet_link_tag "invoice" %>
</head>
<body>
    <%= yield %>
</body>
</html>
```

这里的help函数`wicked_pdf_stylesheet_link_tag`会添加css文件 `app/assets/stylesheets/invoice.scss`

## 修改show.html.erb

参照[这里](https://github.com/PDFTron/rails-generate-pdf/blob/master/app/views/invoices/show.html.erb)

## 添加css文件

修改`config/initializers/asset.rb`

```ruby
# file: config/initializers/assets.rb
Rails.application.config.assets.precompile += %w( invoice.scss )
```

你需要的css和js文件都需要在这里添加

## 添加下载pdf按钮

```ruby
<%= link_to 'DOWNLOAD PDF', invoice_path(@invoice, format: :pdf) %>
```



如果没有问题,就可以看到生成的invoice的pdf文件了

![image-20200309235314882](/../../../../../../../media/2020-03-09-generate-pdf-from-html-in-rails/image-20200309235314882.png)

## 添加footer信息

比较麻烦的地方是footer信息也能使用html来定制样式,根据数据动态生成内容,并且能打印页数.

比较全的配置文件[参考这里](https://github.com/mileszs/wicked_pdf#advanced-usage-with-all-available-options)



