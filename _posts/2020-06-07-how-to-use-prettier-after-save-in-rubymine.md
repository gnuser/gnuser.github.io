---
key: how-to-use-prettier-after-save-in-rubymine
title: How to use prettier after save in RubyMine
tags: RubyMine prettier
date: 2020-06-07 18:26:13 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

如何安装prettier，保存时自动执行，像 `gofmt`一样的工具

<!--more-->

### 安装 prettier

```shell
npm install -g prettier 
npm install -g @prettier/plugin-ruby
```

### 命令行测试

```shell
✗ prettier --write app/models/balance_statistic.rb
app/models/balance_statistic.rb 66ms
```

### 集成进 RubyMine

Preferences -> Tools -> External Tools

添加：

- Name: prettier
- Description: prettier
- Program: /usr/local/bin/prettier
- Arguments: --write $FilePathRelativeToProjectRoot$
- Working directory: $ProjectFileDir$

其他不用改

![image-20200607183232508](/../../../../../../../media/2020-06-07-how-to-use-prettier-after-save-in-rubymine/image-20200607183232508.png)

绑定到keyboard shortcut: `ctrl+s`，或者直接覆盖`command+s`，这样所有的 save 操作都换成了`prettier`,也就实现了保存自动格式化。

![image-20200607183559196](/../../../../../../../media/2020-06-07-how-to-use-prettier-after-save-in-rubymine/image-20200607183559196.png)

### 测试

随便打开一个 ruby 文件，任意输入一些格式不规范的语句，然后 save，或者`ctrl+s`

![image-20200607183730602](/../../../../../../../media/2020-06-07-how-to-use-prettier-after-save-in-rubymine/image-20200607183730602.png)

可以看到prettier 的输出。

目前我暂时还是把`prettier`和`save`分开的，一个是`ctrl+s`，一个是`command+s`，感觉没必要每次都执行一下 prettier,以免被干扰，其实个人认为`git commit` 的时候再自动执行`prettier`，因为现在的 IDE 和个人习惯一般都不会写出太丑的代码，除非遇到一些长语句用一下，正常情况也不需要。