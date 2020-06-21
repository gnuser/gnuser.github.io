---
title: lsof-查看指定进程打开了哪些文件
permalink: /lpo/io/lsof
key: io-lsof
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

lsof 查看指定进程打开了哪些文件

<!--more-->

```shell
$ lsof -p 18940
COMMAND   PID USER   FD   TYPE DEVICE  SIZE/OFF    NODE NAME
python  18940 root  cwd    DIR   0,50      4096 1549389 /
python  18940 root  rtd    DIR   0,50      4096 1549389 /
…
python  18940 root    2u   CHR  136,0       0t0       3 /dev/pts/0
python  18940 root    3w   REG    8,1 117944320     303 /tmp/logtest.txt
```

需要关注的数据列：

- FD 表示文件描述符号
- TYPE 表示文件类型
- NAME 表示文件路径
