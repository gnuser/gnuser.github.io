---
title: dig-递归查询DNS解析过程
permalink: /lpo/network/dig
key: network-dig
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

<!--more-->

```shell
 ~ dig +trace +nodnssec time.geekbang.org

; <<>> DiG 9.10.6 <<>> +trace +nodnssec time.geekbang.org
;; global options: +cmd
;; 第一部分，是查到的一些根域名服务器（.）的 NS 记录。
.			342811	IN	NS	f.root-servers.net.
.			342811	IN	NS	g.root-servers.net.
.			342811	IN	NS	h.root-servers.net.
.			342811	IN	NS	d.root-servers.net.
.			342811	IN	NS	i.root-servers.net.
.			342811	IN	NS	m.root-servers.net.
.			342811	IN	NS	a.root-servers.net.
.			342811	IN	NS	k.root-servers.net.
.			342811	IN	NS	j.root-servers.net.
.			342811	IN	NS	e.root-servers.net.
.			342811	IN	NS	l.root-servers.net.
.			342811	IN	NS	c.root-servers.net.
.			342811	IN	NS	b.root-servers.net.
;; Received 811 bytes from 192.168.0.1#53(192.168.0.1) in 216 ms

;; 第二部分，是从 NS 记录结果中选一个（a.root-servers.net），并查询顶级域名 org. 的 NS 记录。
org.			172800	IN	NS	d0.org.afilias-nst.org.
org.			172800	IN	NS	a0.org.afilias-nst.info.
org.			172800	IN	NS	c0.org.afilias-nst.info.
org.			172800	IN	NS	a2.org.afilias-nst.info.
org.			172800	IN	NS	b0.org.afilias-nst.org.
org.			172800	IN	NS	b2.org.afilias-nst.org.
;; Received 448 bytes from 198.41.0.4#53(a.root-servers.net) in 299 ms

;; 第三部分，是从 org. 的 NS 记录中选择一个（d0.org.afilias-nst.org），并查询二级域名 geekbang.org. 的 NS 服务器。
geekbang.org.		86400	IN	NS	dns9.hichina.com.
geekbang.org.		86400	IN	NS	dns10.hichina.com.
;; Received 96 bytes from 199.19.57.1#53(d0.org.afilias-nst.org) in 163 ms

;; 最后一部分，就是从 geekbang.org. 的 NS 服务器（dns9.hichina.com）查询最终主机 time.geekbang.org. 的 A 记录。
time.geekbang.org.	600	IN	A	39.106.233.176
;; Received 62 bytes from 106.11.141.115#53(dns9.hichina.com) in 4 ms
```

- A 记录，用来把域名转换成 IP 地址；
- CNAME 记录，用来创建别名；
- 而 NS 记录，则表示该域名对应的域名服务器地址。

流程图：

```sequence
participant client
participant DNS
participant "a.root-servers.net"
participant "d0.org.afilias-nst.org"
participant "dns9.hichina.com"


client->>DNS: time.geekbang.org
DNS->>"a.root-servers.net": NS .org
"a.root-servers.net"-->>DNS: d0.org.afilias-nst.org
DNS->>"d0.org.afilias-nst.org": NS geekbang.org
"d0.org.afilias-nst.org"-->>DNS: dns9.hichina.com
DNS ->> "dns9.hichina.com": A time.geekbang.org
"dns9.hichina.com" -->> DNS: 39.106.233.176
DNS -->> client: 39.106.233.176
```

- 第一部分，是从 192.168.0.1 查到的一些根域名服务器（.）的 NS 记录。
- 第二部分，是从 NS 记录结果中选一个（a.root-servers.net），并查询顶级域名 org. 的 NS 记录。
- 第三部分，是从 org. 的 NS 记录中选择一个（d0.org.afilias-nst.org），并查询二级域名 geekbang.org. 的 NS 服务器。
- 最后一部分，就是从 geekbang.org. 的 NS 服务器（dns9.hichina.com）查询最终主机 time.geekbang.org. 的 A 记录。