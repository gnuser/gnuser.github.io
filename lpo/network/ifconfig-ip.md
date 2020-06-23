---
title: ifconfig&ip-网络配置
permalink: /lpo/network/ifconfig-ip
key: network/ifconfig
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

<!--more-->

```shell
$ ifconfig eth0
eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 172.18.9.133  netmask 255.255.240.0  broadcast 172.18.15.255
        inet6 fe80::216:3eff:fe0c:a06c  prefixlen 64  scopeid 0x20<link>
        ether 00:16:3e:0c:a0:6c  txqueuelen 1000  (Ethernet)
        RX packets 5821167233  bytes 1768163980547 (1.7 TB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 5409501255  bytes 769240254302 (769.2 GB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
```

```shell
$ ip -s addr show dev eth0
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP group default qlen 1000
    link/ether 00:16:3e:0c:a0:6c brd ff:ff:ff:ff:ff:ff
    inet 172.18.9.133/20 brd 172.18.15.255 scope global dynamic eth0
       valid_lft 315094934sec preferred_lft 315094934sec
    inet6 fe80::216:3eff:fe0c:a06c/64 scope link
       valid_lft forever preferred_lft forever
    RX: bytes  packets  errors  dropped overrun mcast
    1768167508377 5821190511 0       0       0       0
    TX: bytes  packets  errors  dropped carrier collsns
    769243605647 5409524835 0       0       0       0
```

需要关注的几个指标：

- 第一，网络接口的状态标志。ifconfig 输出中的 RUNNING ，或 ip 输出中的 LOWER_UP ，都表示物理网络是连通的，即网卡已经连接到了交换机或者路由器中。如果你看不到它们，通常表示网线被拔掉了
- 第二，MTU 的大小。MTU 默认大小是 1500，根据网络架构的不同（比如是否使用了 VXLAN 等叠加网络），你可能需要调大或者调小 MTU 的数值。
- 第三，网络接口的 IP 地址、子网以及 MAC 地址。这些都是保障网络功能正常工作所必需的，你需要确保配置正确。
- 第四，网络收发的字节数、包数、错误数以及丢包情况，特别是 TX 和 RX 部分的 errors、dropped、overruns、carrier 以及 collisions 等指标不为 0 时，通常表示出现了网络 I/O 问题。其中：
  - errors 表示发生错误的数据包数，比如校验错误、帧同步错误等；
  - dropped 表示丢弃的数据包数，即数据包已经收到了 Ring Buffer，但因为内存不足等原因丢包；
  - overruns 表示超限数据包数，即网络 I/O 速度过快，导致 Ring Buffer 中的数据包来不及处理（队列满）而导致的丢包；
  - carrier 表示发生 carrirer 错误的数据包数，比如双工模式不匹配、物理电缆出现问题等；
  - collisions 表示碰撞数据包数。
