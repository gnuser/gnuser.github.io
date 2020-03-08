---
title: 加速github访问速度
date: 2020-03-08 11:30 +0800
tags: github
---



github有时候真的太慢了,就算用了代理好像也解决不了问题.为了让写blog变得更顺畅,我做了以下尝试.

<!--more-->

## 修改/etc/hosts

去[https://www.ipaddress.com/](https://www.ipaddress.com/)查询`github.com`和`github.global.ssl.fastly.net`的ip地址

修改/etc/hosts文件

```shell
➜  gnuser.github.io git:(master) ✗ sudo vim /etc/hosts
199.232.5.194 github.global.ssl.fastly.net
140.82.114.4 github.com
```

## 刷新DNS缓存

```shell
sudo killall -HUP mDNSResponder
sudo killall mDNSResponderHelper
sudo dscacheutil -flushcache
```

## What, git push报错了

```shell
➜  gnuser.github.io git:(master) ✗ git push
Counting objects: 537, done.
Delta compression using up to 8 threads.
Compressing objects: 100% (507/507), done.
Writing objects: 100% (537/537), 31.93 MiB | 23.63 MiB/s, done.
Total 537 (delta 52), reused 0 (delta 0)
error: RPC failed; curl 55 SSL_write() returned SYSCALL, errno = 32
fatal: The remote end hung up unexpectedly
fatal: The remote end hung up unexpectedly
```

修改git http.postBuffer大小

```
git config http.postBuffer 524288000
```

## 终极大招

上面各种操作我没感觉有太大区别,还是登陆国外的服务器操作,体验飞一般的感觉

1. 打包本地的文件
2. 上传到服务器
3. git pull下blog项目
4. 解压文件,更新blog项目, 我这里比较粗暴,直接删除老的项目内所有的文件,全部替换为新的
5. git commit && git push