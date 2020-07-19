---
title: 协议层
permalink: /lpo/network/protocol-layer
key: network-protocol-layer
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

网络编程最基础的协议部分。

<!--more-->

![img](../media/protocol-layer/f2dbfb5500c2aa7c47de6216ee7098bd.png)

### 七层模型（OSI）

- 应用层，负责为应用程序提供统一的接口。
- 表示层，负责把数据转换成兼容接收系统的格式。
- 会话层，负责维护计算机之间的通信连接。
- 传输层，负责为数据加上传输表头，形成数据包。
- 网络层，负责数据的路由和转发。
- 数据链路层，负责 MAC 寻址、错误侦测和改错。
- 物理层，负责在物理网络中传输数据帧。

### 四层模型（TCP/IP）

- 应用层，负责向用户提供一组应用程序，比如 HTTP、FTP、DNS 等。
- 传输层，负责端到端的通信，比如 TCP、UDP 等。
- 网络层，负责网络包的封装、寻址和路由，比如 IP、ICMP 等。
- 网络接口层，负责网络包在物理网络中的传输，比如 MAC 寻址、错误侦测以及通过网卡传输网络帧等。

### 除了TCP/IP 协议，还有其他类型的协议：

- Firewire
- USB
- Bluetooth
- WiFi



