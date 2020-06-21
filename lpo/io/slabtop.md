---
title: slabtop-查看slab缓存占用排行
permalink: /lpo/io/slabtop
key: io-slabtop
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

内核使用 Slab 机制，管理目录项和索引节点的缓存

<!--more-->

```shell
sudo slabtop
```

```shell
# 按下c按照缓存大小排序，按下a按照活跃对象数排序
$ sudo slabtop
 Active / Total Objects (% used)    : 1063811 / 1112331 (95.6%)
 Active / Total Slabs (% used)      : 32350 / 32350 (100.0%)
 Active / Total Caches (% used)     : 80 / 114 (70.2%)
 Active / Total Size (% used)       : 177303.62K / 197240.19K (89.9%)
 Minimum / Average / Maximum Object : 0.01K / 0.18K / 8.00K

  OBJS ACTIVE  USE OBJ SIZE  SLABS OBJ/SLAB CACHE SIZE NAME
593190 592890   0%    0.10K  15210       39     60840K buffer_head
 19743  10616   0%    1.06K    794       30     25408K ext4_inode_cache
 43400  37074   0%    0.57K   1550       28     24800K radix_tree_node
 88740  88740 100%    0.13K   2958       30     11832K kernfs_node_cache
 15314  15056   0%    0.59K    589       26      9424K inode_cache
 48867  48867 100%    0.19K   2327       21      9308K kmalloc-192
 44289  33525   0%    0.19K   2109       21      8436K dentry
   712    604   0%    7.50K    178        4      5696K task_struct
  7368   6858   0%    0.66K    325       24      5200K proc_inode_cache
 16112  14506   0%    0.20K    848       19      3392K vm_area_struct
  1568   1463   0%    2.00K     98       16      3136K kmalloc-2048
  2864   2725   0%    1.00K    179       16      2864K kmalloc-1024
 10480   8752   0%    0.25K    655       16      2620K filp
```
