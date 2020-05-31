---
key: understand-cpu-avg-load
title: Understand The CPU Avg Load
date: 2020-05-30 23:59:59 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

linux 性能优化读书笔记-02 | 基础篇：到底应该怎么理解“平均负载”？

<!--more-->

## 平均负载

简单来说，平均负载是指单位时间内，系统处于可运行状态和不可中断状态的平均进程数，也就是平均活跃进程数，它和 CPU 使用率并没有直接关系。

- 可运行状态进程: 所谓可运行状态的进程，是指正在使用 CPU 或者正在等待 CPU 的进程，也就是我们常用 ps 命令看到的，处于 R 状态（Running 或 Runnable）的进程。    

- 不可中断状态进程: 不可中断状态的进程则是正处于内核态关键流程中的进程，并且这些流程是不可打断的，比如最常见的是等待硬件设备的 I/O 响应（读写硬盘），也就是我们在 ps 命令中看到的 D 状态（Uninterruptible Sleep，也称为 Disk Sleep）的进程。    

> The **load average** is the sum of the **run queue length** and the number of jobs currently running on the CPUs.
>
> As the authors explain about the Linux kernel, because both of our test processes are CPU-bound they will be in a **TASK_RUNNING** state. This means they are either:
>
> - **running** i.e., currently executing on the CPU
> - **runnable** i.e., waiting in the **run_queue** for the CPU
>
> The **Linux kernel** also checks to see if there are any tasks in a short-term sleep state called **TASK_UNINTERRUPTIBLE**. If there are, they are also included in the load average sample.

### 当平均负载高于 CPU 数量 70% ，检查原因

关于 cpu 数量，可通过命令查看

- grep 'model name' /proc/cpuinfo | wc -l

而cpu 数量的计算过程，我们在群里根据一些资料（[Understanding Linux _proc_cpuinfo.pdf](https://doc.callmematthi.eu/static/webArticles/Understanding%20Linux%20_proc_cpuinfo.pdf), [CPU核心数与线程数](https://zhuanlan.zhihu.com/p/86855590?utm_source=wechat_session&utm_medium=social&utm_oi=27126606594048)），得出一个推算过程：

- 首先使用lscpu 命令

```shell
$ lscpu
Architecture:        x86_64
CPU op-mode(s):      32-bit, 64-bit
Byte Order:          Little Endian
CPU(s):              4
On-line CPU(s) list: 0-3
Thread(s) per core:  1
Core(s) per socket:  4
Socket(s):           1
NUMA node(s):        1
Vendor ID:           GenuineIntel
CPU family:          6
Model:               79
Model name:          Intel(R) Xeon(R) CPU E5-26xx v4
Stepping:            1
CPU MHz:             2394.446
BogoMIPS:            4788.89
Hypervisor vendor:   KVM
Virtualization type: full
L1d cache:           32K
L1i cache:           32K
L2 cache:            4096K
NUMA node0 CPU(s):   0-3
Flags:               fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ss ht syscall nx lm constant_tsc rep_good nopl cpuid tsc_known_freq pni pclmulqdq ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand hypervisor lahf_lm abm 3dnowprefetch pti bmi1 avx2 bmi2 rdseed adx xsaveopt
```



- 推算公式

```
总的逻辑 cpu 数 = 物理 cpu 数 * 每颗物理 cpu 的核心数 * 每个核心的超线程数
cpu count = Socket(s) * Core(s) per socket * Thread(s) per core
# socket 相当于插槽，普通家用机一般就一条，服务器有多条
# 每个 socket 就对应一个物理 cpu 数
# 每颗物理 cpu 可以有多个核
# 多线程和超线程技术是为了提高单个 core 同一时刻能够执行的多线程数的技术（充分利用单个 core 的计算能力，尽量让其“一刻也不得闲”），对应 Thread(s) per core
```

- 源码跟踪

查看 lscpu 的源码，可以看到每个系数的真正来源。

[https://github.com/karelzak/util-linux/blob/2eb527722af2093038bf38d4554c086d20df79c9/sys-utils/lscpu.c#L1081](https://github.com/karelzak/util-linux/blob/2eb527722af2093038bf38d4554c086d20df79c9/sys-utils/lscpu.c#L1081)

[https://github.com/karelzak/util-linux/blob/2eb527722af2093038bf38d4554c086d20df79c9/sys-utils/lscpu.c#L2105](https://github.com/karelzak/util-linux/blob/2eb527722af2093038bf38d4554c086d20df79c9/sys-utils/lscpu.c#L2105)



## uptime，user 数量



```shell
$ uptime
23:52  up 1 day,  1:41, 4 users, load averages: 2.76 2.69 2.60
```

- 使用 w 命令可以看到有哪几个 user，第一行也包含了`uptime`的输出内容

```shell
$ w
13:21:19 up 147 days,  2:33,  1 user,  load average: 0.11, 0.15, 0.16
USER     TTY      FROM             LOGIN@   IDLE   JCPU   PCPU WHAT
ubuntu   pts/0    xxx.xxx.xxx.xxx   13:21    1.00s  0.08s  0.00s w
```

- 数据来源是`/proc/loadavg`

```shell
$ cat /proc/loadavg
0.19 0.15 0.11 1/319 12201
```

- 源码跟踪

[https://github.com/torvalds/linux/blob/9cb1fd0efd195590b828b9b865421ad345a4a145/kernel/sched/loadavg.c#L157](https://github.com/torvalds/linux/blob/9cb1fd0efd195590b828b9b865421ad345a4a145/kernel/sched/loadavg.c#L157)

计算平均值的算法为EMA，这种算法的目的主要是“距离目标预测窗口越近，则数据的价值越高，对未来影响越大”

股票的指标也有`EMA`,相对于 `MA`，确实可以让`趋势化更明显`

