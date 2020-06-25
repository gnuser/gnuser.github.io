---
title: pktgen-测试网络性能
permalink: /lpo/network/pktgen
key: network-pktgen
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

Linux 内核自带的高性能网络测试工具 pktgen。pktgen 支持丰富的自定义选项，方便你根据实际需要构造所需网络包，从而更准确地测试出目标服务器的性能。

<!--more-->

## 加载`pktgen`内核模块
```shell
sudo modprobe pktgen
```

在 Linux 系统中，你并不能直接找到 pktgen 命令。因为 pktgen 作为一个内核线程来运行，需要你加载 pktgen 内核模块后，再通过 /proc 文件系统来交互。
pktgen 在每个 CPU 上启动一个内核线程，并可以通过 /proc/net/pktgen 下面的同名文件，跟这些线程交互；而 pgctrl 则主要用来控制这次测试的开启和停止。

```shell
sudo ps -ef | grep pktgen | grep -v grep
root     21193     2  0 22:27 ?        00:00:00 [kpktgend_0]
root     21195     2  0 22:27 ?        00:00:00 [kpktgend_1]
root     21196     2  0 22:27 ?        00:00:00 [kpktgend_2]
root     21197     2  0 22:27 ?        00:00:00 [kpktgend_3]
```

```shell
ls /proc/net/pktgen/
kpktgend_0  kpktgend_1  kpktgend_2  kpktgend_3  pgctrl
```

在使用 pktgen 测试网络性能时，需要先给每个内核线程 kpktgend_X 以及测试网卡配置 pktgen 选项，然后再通过 pgctrl 启动测试。

以发包测试为例，假设发包机器使用的网卡是 eth0，而目标机器的 IP 地址为 172.18.9.134，MAC 地址为 00:16:3e:06:16:8e。

```shell

# 定义一个工具函数，方便后面配置各种测试选项
function pgset() {
    local result
    echo $1 > $PGDEV

    result=`cat $PGDEV | fgrep "Result: OK:"`
    if [ "$result" = "" ]; then
         cat $PGDEV | fgrep Result:
    fi
}

# 为0号线程绑定eth0网卡
PGDEV=/proc/net/pktgen/kpktgend_0
pgset "rem_device_all"   # 清空网卡绑定
pgset "add_device eth0"  # 添加eth0网卡

# 配置eth0网卡的测试选项
PGDEV=/proc/net/pktgen/eth0
pgset "count 1000000"    # 总发包数量
pgset "delay 5000"       # 不同包之间的发送延迟(单位纳秒)
pgset "clone_skb 0"      # SKB包复制
pgset "pkt_size 64"      # 网络包大小
pgset "dst 172.18.9.134" # 目的IP
pgset "dst_mac 00:16:3e:06:16:8e"  # 目的MAC

# 启动测试
PGDEV=/proc/net/pktgen/pgctrl
pgset "start"
```

## 执行测试脚本
```shell
sudo bash ./pktgen.sh
```

## 查看结果
```shell
sudo cat /proc/net/pktgen/eth0
Params: count 1000000  min_pkt_size: 64  max_pkt_size: 64
     frags: 0  delay: 5000  clone_skb: 0  ifname: eth0
     flows: 0 flowlen: 0
     queue_map_min: 0  queue_map_max: 0
     dst_min: 172.18.9.134  dst_max:
     src_min:   src_max:
     src_mac: 00:16:3e:0c:a0:6c dst_mac: 00:16:3e:06:16:8e
     udp_src_min: 9  udp_src_max: 9  udp_dst_min: 9  udp_dst_max: 9
     src_mac_count: 0  dst_mac_count: 0
     Flags:
Current:
     pkts-sofar: 1000000  errors: 0
     started: 19092544336018us  stopped: 19092549339700us idle: 4273854us
     seq_num: 1000001  cur_dst_mac_offset: 0  cur_src_mac_offset: 0
     cur_saddr: 172.18.9.133  cur_daddr: 172.18.9.134
     cur_udp_dst: 9  cur_udp_src: 9
     cur_queue_map: 0
     flows: 0
Result: OK: 5003681(c729827+d4273854) usec, 1000000 (64byte,0frags)
  199852pps 102Mb/sec (102324224bps) errors: 0
```

测试报告主要分为三个部分：
- 第一部分的 Params 是测试选项；
- 第二部分的 Current 是测试进度，其中， packts so far（pkts-sofar）表示已经发送了 100 万个包，也就表明测试已完成。
- 第三部分的 Result 是测试结果，包含测试所用时间、网络包数量和分片、PPS、吞吐量以及错误数。

这里的结果，吞吐量为102 Mb/s, PPS 为20万

假如是千兆网卡，PPS大约应该为1000Mbps/((64+20)*8bit) = 1.5 Mpps（其中，20B 为以太网帧前导和帧间距的大小）。