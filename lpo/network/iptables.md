---
title: iptables-防火墙
permalink: /lpo/network/iptables
key: network-iptables
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

<!--more-->

### 禁止指定 IP 数据包

```shell
$ iptables -I INPUT -s 192.168.0.2 -p tcp -j REJECT
```

### 限制 syn 并发数

```shell
# 限制syn并发数为每秒1次
$ iptables -A INPUT -p tcp --syn -m limit --limit 1/s -j ACCEPT

# 限制单个IP在60秒新建立的连接数为10
$ iptables -I INPUT -p tcp --dport 80 --syn -m recent --name SYN_FLOOD --update --seconds 60 --hitcount 10 -j REJECT
```

### 根据 NAT 原理，使用 NAPT方式，将内网 IP映射到公网 IP的不同端口

![img](../media/iptables/c743105dc7bd955a4a300d6b55b7a0e4.png)

假设你现在的机器的内网 IP 为192.168.0.2，路由器分配到的公网 IP为 100.100.100.100

- 当你访问 baidu.com 时，NAT 网关会把源地址，从内网 IP 192.168.0.2 替换成公网 IP 地址 100.100.100.100，然后才发送给 baidu.com；
- 当 baidu.com 发回响应包时，NAT 网关又会把目的地址，从公网 IP 地址 100.100.100.100 替换成内网 IP 192.168.0.2，然后再发送给你。

1. 配置 SNAT，也就是上面的第一步：内网IP -> 公网IP

```shell
$ iptables -t nat -A POSTROUTING -s 192.168.0.0/16 -j MASQUERADE
# 或者
$ iptables -t nat -A POSTROUTING -s 192.168.0.2 -j SNAT --to-source 100.100.100.100
```

这里的POSTROUTING，用于路由判断后所执行的规则，比如，对发送或转发的数据包进行 SNAT 或 MASQUERADE。

2. 配置 DNAT，公网 IP -> 内网 IP

```shell
$ iptables -t nat -A PREROUTING -d 100.100.100.100 -j DNAT --to-destination 192.168.0.2
```

PREROUTING，用于路由判断前所执行的规则，比如，对接收到的数据包进行 DNAT。

3. 双向地址转换,同时添加 SNAT,DNAT 规则

```shell
$ iptables -t nat -A POSTROUTING -s 192.168.0.2 -j SNAT --to-source 100.100.100.100
$ iptables -t nat -A PREROUTING -d 100.100.100.100 -j DNAT --to-destination 192.168.0.2
```

4. 开启IP 转发功能

```shell
$ sysctl -w net.ipv4.ip_forward
net.ipv4.ip_forward = 1
```

