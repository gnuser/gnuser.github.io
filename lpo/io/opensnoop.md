---
title: opensnoop-查看完整文件IO路径
permalink: /lpo/io/opensnoop
key: io-opensnoop
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

同属于 bcc 软件包，可以动态跟踪内核中的 open 系统调用。这样，我们就可以找出这些 txt 文件的路径。

<!--more-->

```shell
$ sudo /usr/share/bcc/tools/opensnoop
PID    COMM               FD ERR PATH
9763   AliYunDun          27   0 /var/log/auth.log
31066  opensnoop          -1   2 /usr/lib/python2.7/encodings/ascii.x86_64-linux-gnu.so
31066  opensnoop          -1   2 /usr/lib/python2.7/encodings/ascii.so
31066  opensnoop          -1   2 /usr/lib/python2.7/encodings/asciimodule.so
31066  opensnoop          14   0 /usr/lib/python2.7/encodings/ascii.py
31066  opensnoop          15   0 /usr/lib/python2.7/encodings/ascii.pyc
9763   AliYunDun          27   0 /var/log/auth.log
9763   AliYunDun          27   0 /proc/9763/stat
9763   AliYunDun          27   0 /sys/devices/system/cpu
9763   AliYunDun          27   0 /proc/9763/stat
3468   ftdc               35   0 /proc/3468/stat
3468   ftdc               35   0 /proc/3468/stat
9763   AliYunDun          27   0 /proc
3468   ftdc               35   0 /proc/stat
3468   ftdc               35   0 /proc/meminfo
3468   ftdc               35   0 /proc/diskstats
3468   mongod             35   0 /var/lib/mongodb/journal
3468   mongod             35   0 /var/lib/mongodb/journal
3468   WTCheck.tThread    35   0 /var/lib/mongodb/WiredTiger.turtle
9763   AliYunDun          27   0 /var/log/auth.log
17654  bitcoin-msghand    98   0 /home/ubuntu/.vds/bitcoin/chainstate/2651917.ldb
17654  bitcoin-msghand    98   0 /home/ubuntu/.vds/bitcoin/chainstate/2651728.ldb
17654  bitcoin-msghand    98   0 /home/ubuntu/.vds/bitcoin/chainstate/2651555.ldb
...
```
