---
title: IO栈全景图
permalink: /lpo/io/stack
key: io-stack
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

根据这张 I/O 栈的全景图，我们可以更清楚地理解，存储系统 I/O 的工作原理。

<!--more-->

![img](../media/linux-storage-stack/14bc3d26efe093d3eada173f869146b1.png)

- 文件系统层，包括虚拟文件系统和其他各种文件系统的具体实现。它为上层的应用程序，提供标准的文件访问接口；对下会通过通用块层，来存储和管理磁盘数据。
- 通用块层，包括块设备 I/O 队列和 I/O 调度器。它会对文件系统的 I/O 请求进行排队，再通过重新排序和请求合并，然后才要发送给下一级的设备层。
- 设备层，包括存储设备和相应的驱动程序，负责最终物理设备的 I/O 操作。
