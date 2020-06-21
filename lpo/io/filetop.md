---
title: filetop-文件IO排行
permalink: /lpo/io/filetop
key: io-iotop
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

加上-C 选项可以滚屏跟踪文件 IO 读写排行

<!--more-->

```shell
# 切换到工具目录
$ cd /usr/share/bcc/tools

# -C 选项表示输出新内容时不清空屏幕
$ ./filetop -C

TID    COMM             READS  WRITES R_Kb    W_Kb    T FILE
514    python           0      1      0       2832    R 669.txt
514    python           0      1      0       2490    R 667.txt
514    python           0      1      0       2685    R 671.txt
514    python           0      1      0       2392    R 670.txt
514    python           0      1      0       2050    R 672.txt

...

TID    COMM             READS  WRITES R_Kb    W_Kb    T FILE
514    python           2      0      5957    0       R 651.txt
514    python           2      0      5371    0       R 112.txt
514    python           2      0      4785    0       R 861.txt
514    python           2      0      4736    0       R 213.txt
514    python           2      0      4443    0       R 45.txt
```

filetop 输出了 8 列内容，分别是：

- 线程 ID
- 线程命令行
- 读写次数
- 读写的大小（单位 KB）
- 文件类型以及读写的文件名称
