---
key: linux-perf-optimization-learning-first-week-summary
title: Linux perf optimization learning first week summary
date: 2020-05-31 23:21:25 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

Linux 性能优化学习，第一周总结

<!--more-->

工具的学习：

- vmstat: 查看上下文切换次数
- pidstat: 查看上下文切换次数
- sysbench: 模拟高CPU占用率，使用多线程
- perf：性能分析工具，可以生成报告，带树形图，使用方法为 `perf record -g && perf report`，也可以指定进程 id`perf top -g -p 21515`
- sar: 查看每秒生成线程或者进程数量
- ptree: 查看进程树
- execsnoop: 专为短时进程设计的工具
- ab: 压测工具
- dstat: 全指标分析工具
- strace: 最常用的跟踪进程系统调用的工具

遇到的问题：

- 如何控制 CPU 画正弦函数？
- 测试环境准备还是不够充分
- 同时进行的还有 ruby 框架 sequent 的深入学习，可能不应该同时进行几个大型的任务，太难了

好的学习方式：

- 提出问题，在一起讨论，就算没得到最终答案，至少印象会非常深刻，这是最好的学习方式

