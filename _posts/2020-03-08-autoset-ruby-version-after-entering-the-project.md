---
key: autoset-ruby-version-after-entering-the-project
title: 进入目录时自动设置ruby版本
tags: jekyll ruby
typora-root-url: "/Users/chenjing/workspace/github/gnuser.github.io"
---



使用rvm,可以管理多个ruby版本,以适应各种项目的ruby版本需求,每次进入项目时,就可以手动指定ruby版本,像这样:

<!--more-->

```shell
$ rvm use 2.6.5
```

为了不用每次都输入这个命令,参考[rvm官网文档](https://rvm.io/workflow/projects)以及这篇博客[^1]. 我们不仅给项目指定ruby版本,还为项目生成一个独立的gem管理目录. 



我当前用的ruby版本是`2.6.5`,博客路径在`/Users/chenjing/workspace/github/gnuser.github.io`.



## 创建gemset

```shell
✗ rvm gemset create gnuser_blog
ruby-2.6.5 - #gemset created /Users/chenjing/.rvm/gems/ruby-2.6.5@gnuser_blog
```

## 指定ruby版本

```shell
✗ rvm --rvmrc ruby-2.6.5@gnuser_blog
Using /Users/chenjing/.rvm/gems/ruby-2.6.5 with gemset gnuser_blog
```

## 确认配置

```shell
gnuser.github.io git:(master) ✗ cd ../
➜  github cd gnuser.github.io
You are using '.rvmrc', it requires trusting, it is slower and it is not compatible with other ruby managers,
you can switch to '.ruby-version' using 'rvm rvmrc to ruby-version'
or ignore this warning with 'rvm rvmrc warning ignore /Users/chenjing/workspace/github/gnuser.github.io/.rvmrc',
'.rvmrc' will continue to be the default project file in RVM 1 and RVM 2,
to ignore the warning for all files run 'rvm rvmrc warning ignore all.rvmrcs'.

*****************************************************************************************************
* NOTICE                                                                                            *
*****************************************************************************************************
* RVM has encountered a new or modified .rvmrc file in the current directory, this is a shell       *
* script and therefore may contain any shell commands.                                              *
*                                                                                                   *
* Examine the contents of this file carefully to be sure the contents are safe before trusting it!  *
* Do you wish to trust '/Users/chenjing/workspace/github/gnuser.github.io/.rvmrc'?                  *
* Choose v[iew] below to view the contents                                                          *
*****************************************************************************************************
y[es], n[o], v[iew], c[ancel]>y
➜  gnuser.github.io git:(master) ✗ bundle install
```

## 确认生效

```shell
➜  gnuser.github.io git:(master) ✗ rvm list
   ruby-2.1.1 [ x86_64 ]
   ruby-2.2.0 [ x86_64 ]
   ruby-2.2.2 [ x86_64 ]
 * ruby-2.3.1 [ x86_64 ]
   ruby-2.3.8 [ x86_64 ]
   ruby-2.4.1 [ x86_64 ]
   ruby-2.5.1 [ x86_64 ]
   ruby-2.5.3 [ x86_64 ]
   ruby-2.6.3 [ x86_64 ]
=> ruby-2.6.5 [ x86_64 ]

# => - current
# =* - current && default
#  * - default
```

`=>`指向的就是当前的ruby版本

## .rvmrc加进.gitignore




[^1]: https://coderwall.com/p/bvkgtw/easy-way-to-create-gemset-per-project-using-rvm