---
title: df-查看磁盘使用情况
permalink: /lpo/io/df
key: io-df
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

查看文件系统磁盘空间使用情况

<!--more-->

## 查看文件系统磁盘的使用情况

```shell
df -h
```

```shell
$ df -h
Filesystem      Size  Used Avail Use% Mounted on
udev            3.9G     0  3.9G   0% /dev
tmpfs           798M  5.9M  792M   1% /run
/dev/vda1       394G  355G   23G  95% /
tmpfs           3.9G     0  3.9G   0% /dev/shm
tmpfs           5.0M     0  5.0M   0% /run/lock
tmpfs           3.9G     0  3.9G   0% /sys/fs/cgroup
tmpfs           798M     0  798M   0% /run/user/0
tmpfs           798M     0  798M   0% /run/user/1000
```

## 查看索引节点的使用情况

```shell
df -i
```

```shell
$ df -i
Filesystem       Inodes  IUsed    IFree IUse% Mounted on
udev            1015447    402  1015045    1% /dev
tmpfs           1020983   1817  1019166    1% /run
/dev/vda1      26214400 394935 25819465    2% /
tmpfs           1020983      1  1020982    1% /dev/shm
tmpfs           1020983      2  1020981    1% /run/lock
tmpfs           1020983     18  1020965    1% /sys/fs/cgroup
tmpfs           1020983     14  1020969    1% /run/user/0
tmpfs           1020983     10  1020973    1% /run/user/1000
```
