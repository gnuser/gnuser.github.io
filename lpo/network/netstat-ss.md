---
title: netstat&ss-套接字信息
permalink: /lpo/network/netstat-ss
key: network/ifconfig-ss
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

可以用 netstat 或者 ss ，来查看套接字、网络栈、网络接口以及路由表的信息。

<!--more-->

```shell
# head -n 3 表示只显示前面3行
# -l 表示只显示监听套接字
# -n 表示显示数字地址和端口(而不是名字)
# -p 表示显示进程信息
$ netstat -nlp | head -n 3
(Not all processes could be identified, non-owned process info
 will not be shown, you would have to be root to see it all.)
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name
tcp        0      0 127.0.0.1:8333          0.0.0.0:*               LISTEN      17654/vdsd
```

```shell
# -l 表示只显示监听套接字
# -t 表示只显示 TCP 套接字
# -n 表示显示数字地址和端口(而不是名字)
# -p 表示显示进程信息
$ ss -ltnp | head -n 3
State    Recv-Q    Send-Q        Local Address:Port        Peer Address:Port
LISTEN   0         128               127.0.0.1:8333             0.0.0.0:*        users:(("vdsd",pid=17654,fd=100))
LISTEN   0         128                 0.0.0.0:80               0.0.0.0:*
```

当套接字处于连接状态（Established）时

- Recv-Q 表示套接字缓冲还没有被应用程序取走的字节数（即接收队列长度）。
- 而 Send-Q 表示还没有被远端主机确认的字节数（即发送队列长度）。

当套接字处于监听状态（Listening）时

- Recv-Q 表示全连接队列的长度。
- 而 Send-Q 表示全连接队列的最大长度。

## 查看协议栈信息

```shell
$ netstat -s
...
Tcp:
    1393561 active connection openings
    821302 passive connection openings
    193019 failed connection attempts
    35969 connection resets received
    22 connections established
    5019550354 segments received
    4754061729 segments sent out
    73038616 segments retransmitted
    4844 bad segments received
    757515 resets sent
    InCsumErrors: 1215
```

```shell
$ ss -s
Total: 490 (kernel 0)
TCP:   39 (estab 22, closed 2, orphaned 0, synrecv 0, timewait 2/0), ports 0

Transport Total     IP        IPv6
*	  0         -         -
RAW	  1         0         1
UDP	  253       252       1
TCP	  37        34        3
INET	  291       286       5
FRAG	  0         0         0
```

- ss 只显示已经连接、关闭、孤儿套接字等简要统计
- netstat 则提供的是更详细的网络协议栈信息
