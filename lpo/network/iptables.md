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
