---
key: vnpy-cannot-click-menu
title: vnpy-cannot-click-menu
date: 2020-12-24 23:29:20 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---
解决mac下vnpy菜单栏无法点击的问题
<!--more-->

## 修改vnpy/trader/ui/mainwindow.py

添加`bar.setNativeMenuBar(False)`

![image-20201224234706150](/../../../../../../../media/2020-12-24-vnpy-cannot-click-menu/image-20201224234706150.png)



效果如下,菜单栏回到到应用窗口上

![image-20201224234742417](/../../../../../../../media/2020-12-24-vnpy-cannot-click-menu/image-20201224234742417.png)