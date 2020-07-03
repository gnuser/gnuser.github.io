---
title: hping3-DDOS之SYN攻击
permalink: /lpo/network/hping3
key: network-hping3
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

使用 hping3 来模拟 DDOS 攻击

<!--more-->

```shell
# -S参数表示设置TCP协议的SYN（同步序列号），-p表示目的端口为80
# -i u10表示每隔10微秒发送一个网络帧
$ hping3 -S -p 80 -i u10 192.168.0.30
```

- 如果现象不明显，可以减小 u10（比如改为 u1）,或者添加-flood 选项
- 可以添加--rand-source 随机化源 IP

## 原理

- 即客户端构造大量的 SYN 包，请求建立 TCP 连接；
- 而服务器收到包后，会向源 IP 发送 SYN+ACK 报文，并等待三次握手的最后一次 ACK 报文，直到超时。
- 这种等待状态的 TCP 连接，通常也称为半开连接。由于连接表的大小有限，大量的半开连接就会导致连接表迅速占满，从而无法建立新的 TCP 连接。

## 解决方法

- 使用 [iptables](/lpo/network/iptables) 限制网络包
- 调整系统参数

1. 半开连接数

```shell
# 查看配置
$ sysctl net.ipv4.tcp_max_syn_backlog
net.ipv4.tcp_max_syn_backlog = 256
# 修改配置
$ sysctl -w net.ipv4.tcp_max_syn_backlog=1024
net.ipv4.tcp_max_syn_backlog = 1024
```

2. SYN_RECV 重试次数

```shell
$ sysctl -w net.ipv4.tcp_synack_retries=1
net.ipv4.tcp_synack_retries = 1
```

3. 开启 TCP SYN Cookies

```shell
$ sysctl -w net.ipv4.tcp_syncookies=1
net.ipv4.tcp_syncookies = 1
```

上述配置重启后会丢失，如果需要永久保存，写入`/etc/sysctl.conf`

```shell
$ cat /etc/sysctl.conf
net.ipv4.tcp_syncookies = 1
net.ipv4.tcp_synack_retries = 1
net.ipv4.tcp_max_syn_backlog = 1024
```

## 测试网络延迟

```shell
# -c表示发送3次请求，-S表示设置TCP SYN，-p表示端口号为80
$ hping3 -c 3 -S -p 80 baidu.com
HPING baidu.com (eth0 220.181.38.148): S set, 40 headers + 0 data bytes
len=40 ip=220.181.38.148 ttl=250 id=2920 sport=80 flags=SA seq=0 win=8192 rtt=43.9 ms
len=40 ip=220.181.38.148 ttl=250 id=17036 sport=80 flags=SA seq=1 win=8192 rtt=39.8 ms
len=40 ip=220.181.38.148 ttl=250 id=38340 sport=80 flags=SA seq=2 win=8192 rtt=39.8 ms

--- baidu.com hping statistic ---
3 packets transmitted, 3 packets received, 0% packet loss
round-trip min/avg/max = 39.8/41.1/43.9 ms
```

`rtt`为网络延迟，这里差不多 40ms
