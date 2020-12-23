---
key: rails-may-have-been-in-progress-in-another-thread-when-fork-was-called
title: rails-may-have-been-in-progress-in-another-thread-when-fork-was-called
date: 2020-12-23 17:25:46 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

rails c时碰到一个奇怪的错误 may have been in progress in another thread when fork() was called
<!--more-->

解决方法：

把sprint stop重启，或者启动前设置环境变量`export DISABLE_SPRING=true`

