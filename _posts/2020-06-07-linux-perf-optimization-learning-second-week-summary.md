---
key: linux-perf-optimization-learning-second-week-summary
title: Linux perf optimization learning second week summary
date: 2020-06-07 10:44:21 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

Linux 性能优化学习，第二周总结

<!--more-->

发现作者的总结很完善，还整理成了表格

![img](/../../../../../../../media/2020-06-07-linux-perf-optimization-learning-second-week-summary/596397e1d6335d2990f70427ad4b14ec.png)

常用的 3 大工具针对的指标：

![img](/../../../../../../../media/2020-06-07-linux-perf-optimization-learning-second-week-summary/7a445960a4bc0a58a02e1bc75648aa17.png)

## 拿几个工具讲一下使用心得

### atop

以前没怎么用过`atop`，使用了一下，感觉面板非常全面，`CPU,MEM,NETIO,DISKIO`都有很好的展示

![image-20200607105001055](/../../../../../../../media/2020-06-07-linux-perf-optimization-learning-second-week-summary/image-20200607105001055.png)

可以看到上图中，`MEM`, `SWP` 都是红的，8G 的内存基本耗尽了，按`m`键可以按`MEM`占用排序，按`i`键后可以设置刷新时间，默认是`10s`，可以设置成`3`s

### dtat

这个滚动显示的工具其实很强，可以全面显示系统资源情况

![31C18656-C30D-4661-9110-2E78F34BC825](/../../../../../../../media/2020-06-07-linux-perf-optimization-learning-second-week-summary/31C18656-C30D-4661-9110-2E78F34BC825.png)

这是以前做性能压测的时候显示的信息，命令是`dstat -tfvnrl`，可以看到4 个`cpu`都超过 70%，`int,csw`也比较高`8000+`

### perf

`perf record -g` + `perf report` 这个组合非常好用，能看到占用高 CPU进程的调用堆栈，但是对没有符号的进程是看不到的

![image-20200607114828654](/../../../../../../../media/2020-06-07-linux-perf-optimization-learning-second-week-summary/image-20200607114828654.png)

比如这里的 vdsd 进程，没法展开了，所以在编译程序的时候最好是能加上调试符号信息，只是增加一点文件大小，但对后续调试分析会有很大帮助，当然相反如果你不想让别人分析你的程序，那么你就把调试符号干掉。

### 一个CPU 100%解决的真实案例

1. 使用 htop 发现 cpu 100%
2. 查看到进程是`puma`几乎占满了 cpu
3. 使用 `strace -p pid` ，发现获取不到什么log
4. 使用ps -efL | grep puma, 获取到一堆的线程

```shell
ubuntu    4549 16342  4549  0    9 May19 ?        00:00:00 puma: cluster worker 0: 16342 [payment]
ubuntu    4549 16342  4551  0    9 May19 ?        00:00:00 puma: cluster worker 0: 16342 [payment]
ubuntu    4549 16342  4558  0    9 May19 ?        00:00:00 puma: cluster worker 0: 16342 [payment]
ubuntu    4549 16342  4585  0    9 May19 ?        00:00:29 puma: cluster worker 0: 16342 [payment]
ubuntu    4549 16342  4587  0    9 May19 ?        00:00:03 puma: cluster worker 0: 16342 [payment]
ubuntu    4549 16342  4588  0    9 May19 ?        00:01:10 puma: cluster worker 0: 16342 [payment]
ubuntu    4549 16342  4589  0    9 May19 ?        00:00:01 puma: cluster worker 0: 16342 [payment]
ubuntu    4549 16342  4590  0    9 May19 ?        00:00:00 puma: cluster worker 0: 16342 [payment]
ubuntu    4549 16342  8450  0    9 May24 ?        00:00:03 puma: cluster worker 0: 16342 [payment]
```

这里的第 2，3，4 列分别代表 `ppid pid threadid`

5. 再使用strace -p threadId, 发现一堆的网络请求, 一些数据库查询请求操作关键字
6. 然后使用tcpdump抓包,获取cap文件,再用 wireshark 打开发现确实是一堆的查询

```shell
tcpdump tcp port 9997 -vvv -nnn -s0 -w qxt_oqc.cap
```

7. 交给后台程序员,发现是个数据库操作,异常后会retry,然后就无限retry
8. 改掉重启，问题解决



