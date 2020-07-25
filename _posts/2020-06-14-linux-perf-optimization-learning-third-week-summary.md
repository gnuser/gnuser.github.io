---
key: linux-perf-optimization-learning-third-week-summary
title: Linux性能学习第三周-内存性能优化
date: 2020-06-14 20:22:07 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
sidebar:
  nav: lpo-note
permalink: /lpo-note/3rd-week
---

Linux 性能优化学习，第三周总结

本周主要是内存性能优化

<!--more-->

### free命令的free 不是真正的可用内存

```shell
$ free -h
              total        used        free      shared  buff/cache   available
Mem:           7.6G        5.6G        210M        1.8M        1.8G        1.7G
Swap:          8.0G        1.8G        6.2G
```

free 是未使用内存的大小，这里只有 210M

available 是新进程可用内存的大小，这里有 1.7G， available 不仅包含未使用内存，还包括了可回收的缓存，所以一般会比未使用内存更大。

buffer 和 cache 分别缓存磁盘和文件系统的读写数据

### 如何统计所有进程的物理内存使用量

```shell
$ sudo grep Pss /proc/[1-9]*/smaps | awk '{total+=$2}; END {printf "%d kB\n", total }'
7529991 kB
```



### 使用 glances工具

这个工具也比较全面，还可以显示外网 IP

![image-20200614213020605](/../../../../../../../media/2020-06-14-linux-perf-optimization-learning-third-week-summary/image-20200614213020605.png)

可以看到右上角关于内存的使用情况: 

- total: 7.63G
- used: 6.02G
- free: 1.62G
- active: 5.09G
- inactive: 1.96G
- buffers: 74.2M
- cached: 1.64G
- SWAP 使用了22%

### 使用 memleak 检查运行中的进程是否有内存泄漏

```shell
$ sudo /usr/share/bcc/tools/memleak -p $(pidof vdsd) -a
	addr = 7fcf8e2111d0 size = 65536
	addr = 7fd0e03e6fa0 size = 65536
	addr = 7fcf8e1e11a0 size = 65536
	addr = 7fcf8e1f11b0 size = 65536
	addr = 7fcf8e2011c0 size = 65536
	addr = 7fd0a0b12560 size = 444816
	addr = 7fd0a0caa880 size = 758248
	7160 bytes in 1 allocations from stack
		operator new(unsigned long)+0x18 [libstdc++.so.6.0.25]
		[unknown]
		[unknown]
	7160 bytes in 1 allocations from stack
		[unknown] [vdsd]
		[unknown]
	10880 bytes in 272 allocations from stack
		operator new(unsigned long)+0x18 [libstdc++.so.6.0.25]
		[unknown]
	17143 bytes in 337 allocations from stack
		[unknown] [vdsd]
	20336 bytes in 82 allocations from stack
		operator new(unsigned long)+0x18 [libstdc++.so.6.0.25]
		[unknown] [vdsd]
		[unknown] [vdsd]
		[unknown] [vdsd]
		[unknown]
	38000 bytes in 250 allocations from stack
		operator new(unsigned long)+0x18 [libstdc++.so.6.0.25]
		[unknown]
	56948 bytes in 845 allocations from stack
		operator new(unsigned long)+0x18 [libstdc++.so.6.0.25]
	78736 bytes in 703 allocations from stack
		operator new(unsigned long)+0x18 [libstdc++.so.6.0.25]
		[unknown]
		[unknown] [vdsd]
		[unknown]
	536192 bytes in 10 allocations from stack
		operator new(unsigned long)+0x18 [libstdc++.so.6.0.25]
		[unknown]
		[unknown]
	1433130 bytes in 1158 allocations from stack
		operator new(unsigned long)+0x18 [libstdc++.so.6.0.25]
```

这里通过检查 vdsd 进程，省略了一些输出，发现有好多内存分配的调用栈，但由于没有符号，没办法确认泄漏的代码

### 再使用 pmap 检查进程内存分布

```shell
$ sudo pmap -x $(pidof vdsd) | less
Address           Kbytes     RSS   Dirty Mode  Mapping
000055ef3f9cd000   31680    7872       8 r-x-- vdsd
000055ef3f9cd000       0       0       0 r-x-- vdsd
000055ef41abc000     560     560     560 r---- vdsd
000055ef41abc000       0       0       0 r---- vdsd
000055ef41b48000     124      96      96 rw--- vdsd
000055ef41b48000       0       0       0 rw--- vdsd
000055ef41b67000     472     472     472 rw---   [ anon ]
000055ef41b67000       0       0       0 rw---   [ anon ]
000055ef41bdd000       4       4       4 rw---   [ anon ]
000055ef41bdd000       0       0       0 rw---   [ anon ]
000055ef41bde000     340     332     332 rw---   [ anon ]
000055ef41bde000       0       0       0 rw---   [ anon ]
000055ef431c1000   10364   10304   10304 rw---   [ anon ]
000055ef431c1000       0       0       0 rw---   [ anon ]
000055ef43be0000       4       4       4 rw---   [ anon ]
000055ef43be0000       0       0       0 rw---   [ anon ]
000055ef43be1000    6812    6796    6796 rw---   [ anon ]
000055ef43be1000       0       0       0 rw---   [ anon ]
000055ef44288000       4       4       4 rw---   [ anon ]
000055ef44288000       0       0       0 rw---   [ anon ]
省略。。。
00007fffffffe000       0       0       0 --x--   [ anon ]
ffffffffff600000       4       0       0 r-x--   [ anon ]
ffffffffff600000       0       0       0 r-x--   [ anon ]
---------------- ------- ------- -------
total kB         6786588 3526248 3494408

```

看到 vdsd 的进程内存分布比较分散，应该还是分配了内存但没回收，但是由于看不到符号，代码也没开源，暂时没有办法进一步的分析，尝试观察一下他的运行参数，降低一些可能的内存消耗，或者限制他的内存最大使用量，不知道这样会不会更容易触发 OOM 杀掉。

### swap 原理

一个很典型的场景就是，即使内存不足时，有些应用程序也并不想被 OOM 杀死，而是希望能缓一段时间，等待人工介入，或者等系统自动释放其他进程的内存，再分配给它。

### 快速定位内存问题

![img](/../../../../../../../media/2020-06-14-linux-perf-optimization-learning-third-week-summary/d79cd017f0c90b84a36e70a3c5dccffe.png)

1. 先用 free 和 top，查看系统整体的内存使用情况。
2. 再用 vmstat 和 pidstat，查看一段时间的趋势，从而判断出内存问题的类型。
3. 最后进行详细分析，比如内存分配分析、缓存 / 缓冲区分析、具体进程的内存使用分析等。

