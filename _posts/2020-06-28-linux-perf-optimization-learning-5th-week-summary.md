---
key: linux-perf-optimization-learning-5th-week-summary
title: Linux perf optimization learning 5th week summary
date: 2020-06-28 19:21:26 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

linux 性能优化学习第五周总结，本周主要是 网络性能优化

<!--more-->

## 网络性能指标：

- **带宽**，表示链路的最大传输速率，单位通常为 b/s （比特 / 秒）。	
- **吞吐量**，表示单位时间内成功传输的数据量，单位通常为 b/s（比特 / 秒）或者 B/s（字节 / 秒）。吞吐量受带宽限制，而吞吐量 / 带宽，也就是该网络的使用率。
- **延时**，表示从网络请求发出后，一直到收到远端响应，所需要的时间延迟。在不同场景中，这一指标可能会有不同含义。比如，它可以表示，建立连接需要的时间（比如 TCP 握手延时），或一个数据包往返所需的时间（比如 RTT）。
- **PPS**，是 Packet Per Second（包 / 秒）的缩写，表示以网络包为单位的传输速率。PPS 通常用来评估网络的转发能力，比如硬件交换机，通常可以达到线性转发（即 PPS 可以达到或者接近理论最大值）。而基于 Linux 服务器的转发，则容易受网络包大小的影响。

除了这些指标，网络的可用性（网络能否正常通信）、并发连接数（TCP 连接数量）、丢包率（丢包百分比）、重传率（重新传输的网络包比例）等也是常用的性能指标。

## C10K

C10K 指支持 1万的并发请求。

1. 硬件预估

- 假如每个请求消耗内存不到 200KB，200KB*10000 = 2G，总的内存需要 2GB
- 千兆网卡，1000MBit / 10000 = 100Kbit， 带宽只要 100Kbit

2. IO 模型

- 从同步阻塞，为每个请求分配一个线程或进程换成异步非阻塞，更少的线程，IO多路复用
  - 使用非阻塞 I/O 和水平触发通知，比如使用 select 或者 poll。
  - 使用非阻塞 I/O 和边缘触发通知，比如 epoll。
  - 使用异步 I/O（Asynchronous I/O，简称为 AIO）。

水平触发和边缘触发：

- 水平触发：只要文件描述符可以非阻塞地执行 I/O ，就会触发通知。也就是说，应用程序可以随时检查文件描述符的状态，然后再根据状态，进行 I/O 操作。
- 边缘触发：只有在文件描述符的状态发生改变（也就是 I/O 请求达到）时，才发送一次通知。这时候，应用程序需要尽可能多地执行 I/O，直到无法继续读写，才可以停止。如果 I/O 没执行完，或者因为某种原因没来得及处理，那么这次通知也就丢失了。

epoll:

- epoll 使用红黑树，在内核中管理文件描述符的集合，这样，就不需要应用程序在每次操作时都传入、传出这个集合。
- epoll 使用事件驱动的机制，只关注有 I/O 事件发生的文件描述符，不需要轮询扫描整个集合。

3. 工作模型优化

- 主进程 + 多个 worker子进程
  - 主进程执行 bind() + listen() 后，创建多个子进程；
  - 然后，在每个子进程中，都通过 accept() 或 epoll_wait() ，来处理相同的套接字。
- 监听到相同端口的多进程模型

## C100K

理论还是基于 C10K，epoll 配合线程池，再加上 CPU、内存和网络接口的性能和容量提升。C100K 很自然就可以达到。

## C1000K

C1000K 的解决方法，本质上还是构建在 epoll 的非阻塞 I/O 模型上。只不过，除了 I/O 模型之外，还需要从应用程序到 Linux 内核、再到 CPU、内存和网络等各个层次的深度优化，特别是需要借助硬件，来卸载那些原来通过软件处理的大量功能。

## C10M

要解决这个问题，最重要就是跳过内核协议栈的冗长路径，把网络包直接送到要处理的应用程序那里去。这里有两种常见的机制，DPDK 和 XDP。

## 问题：如何评估QPS/TPS
