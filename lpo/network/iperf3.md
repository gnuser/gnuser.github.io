---
title: iperf3-TCP/UDP性能
permalink: /lpo/network/iperf3
key: network/iperf3
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

iperf 和 netperf 都是最常用的网络性能测试工具，测试 TCP 和 UDP 的吞吐量。它们都以客户端和服务器通信的方式，测试一段时间内的平均吞吐量。

<!--more-->

## 安装
```shell
# Ubuntu
apt-get install iperf3
# CentOS
yum install iperf3
```

## 测试
1. 在目标机器启动`iperf`服务端
```shell
# -s表示启动服务端，-i表示汇报间隔，-p表示监听端口
iperf3 -s -i 1 -p 8332
```

2. 在另一台机器运行`iperf`客户端
```shell
# -c表示启动客户端，192.168.0.30为目标服务器的IP
# -b表示目标带宽(单位是bits/s)
# -t表示测试时间
# -P表示并发数，-p表示目标服务器监听端口
$ iperf3 -c 192.168.0.30 -b 1G -t 15 -P 2 -p 10000

[  4]  12.00-13.00  sec  69.0 MBytes   579 Mbits/sec  243   35.4 KBytes
[  6]  12.00-13.00  sec  0.00 Bytes  0.00 bits/sec    0   1.41 KBytes
[SUM]  12.00-13.00  sec  69.0 MBytes   579 Mbits/sec  243
- - - - - - - - - - - - - - - - - - - - - - - - -
[  4]  13.00-14.00  sec  57.5 MBytes   482 Mbits/sec   99   43.8 KBytes
[  6]  13.00-14.00  sec  0.00 Bytes  0.00 bits/sec    0   1.41 KBytes
[SUM]  13.00-14.00  sec  57.5 MBytes   482 Mbits/sec   99
- - - - - - - - - - - - - - - - - - - - - - - - -
[  4]  14.00-15.00  sec  60.2 MBytes   505 Mbits/sec  379   19.8 KBytes
[  6]  14.00-15.00  sec  0.00 Bytes  0.00 bits/sec    0   1.41 KBytes
[SUM]  14.00-15.00  sec  60.2 MBytes   505 Mbits/sec  379
- - - - - - - - - - - - - - - - - - - - - - - - -
[ ID] Interval           Transfer     Bandwidth       Retr
[  4]   0.00-15.00  sec   722 MBytes   404 Mbits/sec  4480             sender
[  4]   0.00-15.00  sec   721 MBytes   403 Mbits/sec                  receiver
[  6]   0.00-15.00  sec   217 MBytes   122 Mbits/sec  2456             sender
[  6]   0.00-15.00  sec   216 MBytes   121 Mbits/sec                  receiver
[SUM]   0.00-15.00  sec   939 MBytes   525 Mbits/sec  6936             sender
[SUM]   0.00-15.00  sec   937 MBytes   524 Mbits/sec                  receiver
```

最后的 SUM 行就是测试的汇总结果，包括测试时间、数据传输量以及带宽等。按照发送和接收，这一部分又分为了 sender 和 receiver 两行。
从测试结果你可以看到，这台机器 TCP 接收的带宽（吞吐量）为 524 Mb/s,接收的带宽（吞吐量）为 525 Mb/s
