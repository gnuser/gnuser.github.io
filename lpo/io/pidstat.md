---
title: pidstat-查看进程IO
permalink: /lpo/io/pidstat
key: io-pidstat
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

pidstat 可以观察哪些进程在进行磁盘读写

<!--more-->

```shell
sudo pidstat -d 1
```

```shell
$ sudo pidstat -d 1
Linux 4.15.0-52-generic (iZwz93d1of4sbhantbl6xfZ) 	06/21/20 	_x86_64_	(4 CPU)

19:32:00      UID       PID   kB_rd/s   kB_wr/s kB_ccwr/s iodelay  Command
19:32:01        0      9763      0.00      3.96      0.00       0  AliYunDun

19:32:01      UID       PID   kB_rd/s   kB_wr/s kB_ccwr/s iodelay  Command
19:32:02     1000     17654      0.00      8.00      0.00       0  vdsd

19:32:02      UID       PID   kB_rd/s   kB_wr/s kB_ccwr/s iodelay  Command

19:32:03      UID       PID   kB_rd/s   kB_wr/s kB_ccwr/s iodelay  Command
^C

Average:      UID       PID   kB_rd/s   kB_wr/s kB_ccwr/s iodelay  Command
Average:        0      9763      0.00      1.00      0.00       0  AliYunDun
Average:     1000     17654      0.00      2.00      0.00       0  vdsd
```

- 用户 ID（UID）和进程 ID（PID） 。每秒读取的数据大小（kB_rd/s） ，单位是 KB。
- 每秒发出的写请求数据大小（kB_wr/s） ，单位是 KB。
- 每秒取消的写请求数据大小（kB_ccwr/s） ，单位是 KB。
- 块 I/O 延迟（iodelay），包括等待同步块 I/O 和换入块 I/O 结束的时间，单位是时钟周期。
