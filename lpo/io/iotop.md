---
title: iotop-进程IO排行
permalink: /lpo/io/iotop
key: io-iotop
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

iotop 可以按照

<!--more-->

```shell
sudo iotop
```

```shell
$ sudo iotop
Total DISK READ :       0.00 B/s | Total DISK WRITE :       0.00 B/s
Actual DISK READ:       0.00 B/s | Actual DISK WRITE:       0.00 B/s
  TID  PRIO  USER     DISK READ  DISK WRITE  SWAPIN     IO>    COMMAND
    1 be/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % systemd --system --deserialize 24
    2 be/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [kthreadd]
    4 be/0 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [kworker/0:0H]
    6 be/0 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [mm_percpu_wq]
    7 be/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [ksoftirqd/0]
    8 be/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [rcu_sched]
    9 be/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [rcu_bh]
   10 rt/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [migration/0]
   11 rt/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [watchdog/0]
   12 be/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [cpuhp/0]
   13 be/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [cpuhp/1]
   14 rt/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [watchdog/1]
   15 rt/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [migration/1]
   16 be/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [ksoftirqd/1]
   18 be/0 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [kworker/1:0H]
   19 be/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [cpuhp/2]
   20 rt/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [watchdog/2]
   21 rt/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [migration/2]
   22 be/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [ksoftirqd/2]
   24 be/0 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [kworker/2:0H]
   25 be/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [cpuhp/3]
   26 rt/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [watchdog/3]
   27 rt/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [migration/3]
   28 be/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [ksoftirqd/3]
   30 be/0 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [kworker/3:0H]
   31 be/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [kdevtmpfs]
   32 be/0 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [netns]
   33 be/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [rcu_tasks_kthre]
   34 be/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [kauditd]
   37 be/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [khungtaskd]
   38 be/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [oom_reaper]
   39 be/0 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [writeback]
   40 be/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [kcompactd0]
   41 be/5 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [ksmd]
   42 be/7 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [khugepaged]
   43 be/0 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [crypto]
   44 be/0 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [kintegrityd]
   45 be/0 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [kblockd]
   46 be/0 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [ata_sff]
   47 be/0 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [md]
```

数据列分别为：

- 线程 ID
- I/O 优先级
- 用户
- 每秒读磁盘的大小
- 每秒写磁盘的大小
- 换入 I/O 的时钟百分比
- 等待 I/O 的时钟百分比等
