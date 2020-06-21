---
title: iostat-磁盘性能检测
permalink: /lpo/io/iostat
key: io-iostat
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

iostat 是最常用的磁盘 I/O 性能观测工具，它提供了每个磁盘的使用率、IOPS、吞吐量等各种常见的性能指标，当然，这些指标实际上来自 /proc/diskstats。

<!--more-->

```shell
sudo iostat -d -x 1
```

```shell
 sudo iostat -d -x 1
Linux 4.15.0-52-generic (iZwz93d1of4sbhantbl6xfZ) 	06/21/20 	_x86_64_	(4 CPU)

Device            r/s     w/s     rkB/s     wkB/s   rrqm/s   wrqm/s  %rrqm  %wrqm r_await w_await aqu-sz rareq-sz wareq-sz  svctm  %util
loop0            0.00    0.00      0.00      0.00     0.00     0.00   0.00   0.00    0.00    0.00   0.00     1.60     0.00   0.00   0.00
vda             17.92   11.06   1166.82   1994.73     0.00    19.04   0.01  63.24    1.03    9.44   0.09    65.10   180.30   0.61   1.78

Device            r/s     w/s     rkB/s     wkB/s   rrqm/s   wrqm/s  %rrqm  %wrqm r_await w_await aqu-sz rareq-sz wareq-sz  svctm  %util
loop0            0.00    0.00      0.00      0.00     0.00     0.00   0.00   0.00    0.00    0.00   0.00     0.00     0.00   0.00   0.00
vda              1.00    0.00      8.00      0.00     0.00     0.00   0.00   0.00    0.00    0.00   0.00     8.00     0.00   0.00   0.00

Device            r/s     w/s     rkB/s     wkB/s   rrqm/s   wrqm/s  %rrqm  %wrqm r_await w_await aqu-sz rareq-sz wareq-sz  svctm  %util
loop0            0.00    0.00      0.00      0.00     0.00     0.00   0.00   0.00    0.00    0.00   0.00     0.00     0.00   0.00   0.00
vda              1.00    0.00     68.00      0.00     0.00     0.00   0.00   0.00    0.00    0.00   0.00    68.00     0.00   0.00   0.00
```

- %util ，就是磁盘 I/O 使用率；
- r/s+ w/s ，就是 IOPS；
- rkB/s+wkB/s ，就是吞吐量；
- r_await+w_await ，就是响应时间。



![img](../media/iostat/cff31e715af51c9cb8085ce1bb48318d.png)