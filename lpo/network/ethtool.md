---
title: ethtool-查看网卡信息
permalink: /lpo/network/ethtool
key: network/ethtool
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

ethtool 命令主要用于查询配置网卡参数。比如带宽速度
<!--more-->

用法：ethtool eth0

```shell
$ ethtool eth0
Settings for eth0:
	Supported ports: [ ]
	Supported link modes:   Not reported
	Supported pause frame use: No
	Supports auto-negotiation: No
	Supported FEC modes: Not reported
	Advertised link modes:  Not reported
	Advertised pause frame use: No
	Advertised auto-negotiation: No
	Advertised FEC modes: Not reported
	Speed: Unknown!
	Duplex: Unknown! (255)
	Port: Other
	PHYAD: 0
	Transceiver: internal
	Auto-negotiation: off
Cannot get wake-on-lan settings: Operation not permitted
	Link detected: yes
```

这里的`Speed`显示`Unknown!`,阿里云服务器查不出来

