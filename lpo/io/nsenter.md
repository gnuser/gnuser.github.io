---
title: nsenter-进入容器命名空间运行指定程序
permalink: /lpo/io/nsenter
key: io-nsenter
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

观察进程网络套接字信息

<!--more-->

```shell
# -i表示显示网络套接字信息
$ sudo nsenter --target 8035 --net -- lsof -i
COMMAND     PID            USER   FD   TYPE  DEVICE SIZE/OFF NODE NAME
systemd-n   887 systemd-network   19u  IPv4   17529      0t0  UDP VM-0-3-ubuntu:bootpc
systemd-r   919 systemd-resolve   12u  IPv4   16915      0t0  UDP localhost:domain
systemd-r   919 systemd-resolve   13u  IPv4   16916      0t0  TCP localhost:domain (LISTEN)
netdata    1113         netdata    3u  IPv4   21982      0t0  TCP *:19999 (LISTEN)
netdata    1113         netdata    6u  IPv4   22461      0t0  UDP localhost.localdomain:8125
netdata    1113         netdata    7u  IPv4   22462      0t0  TCP localhost.localdomain:8125 (LISTEN)
sshd       1198            root    3u  IPv4   19915      0t0  TCP *:ssh (LISTEN)
ntpd       1223             ntp   16u  IPv4   19920      0t0  UDP localhost.localdomain:ntp
ntpd       1223             ntp   17u  IPv4   19922      0t0  UDP VM-0-3-ubuntu:ntp
ntpd       1223             ntp   18u  IPv6   19924      0t0  UDP ip6-localhost:ntp
ntpd       1223             ntp   19u  IPv6   19928      0t0  UDP VM-0-3-ubuntu:ntp
redis-ser  1308           redis    6u  IPv4   22002      0t0  TCP localhost.localdomain:6379 (LISTEN)
redis-ser  1308           redis    7u  IPv6   22003      0t0  TCP ip6-localhost:6379 (LISTEN)
redis-ser  1308           redis    8u  IPv4   47600      0t0  TCP localhost.localdomain:6379->localhost.localdomain:42210 (ESTABLISHED)
redis-ser  1308           redis    9u  IPv4  145449      0t0  TCP localhost.localdomain:6379->localhost.localdomain:48172 (ESTABLISHED)
redis-ser  1308           redis   10u  IPv4  140461      0t0  TCP localhost.localdomain:6379->localhost.localdomain:47908 (ESTABLISHED)
redis-ser  1308           redis   11u  IPv4  268333      0t0  TCP localhost.localdomain:6379->localhost.localdomain:55560 (ESTABLISHED)
redis-ser  1308           redis   12u  IPv4 1348246      0t0  TCP localhost.localdomain:6379->localhost.localdomain:50610 (ESTABLISHED)
redis-ser  1308           redis   13u  IPv4 1348262      0t0  TCP localhost.localdomain:6379->localhost.localdomain:50614 (ESTABLISHED)
epmd       1539        rabbitmq    3u  IPv4   20119      0t0  TCP *:epmd (LISTEN)
epmd       1539        rabbitmq    4u  IPv6   20120      0t0  TCP *:epmd (LISTEN)
```

- FD 表示操作的文件描述符
