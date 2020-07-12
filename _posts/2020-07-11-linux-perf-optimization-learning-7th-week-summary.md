---
key: linux-perf-optimization-learning-7th-week-summary
title: Linux perf optimization learning 7th week summary
date: 2020-07-11 12:11:35 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

linux 性能优化学习第七周总结，本周是综合实战

<!--more-->

## 容器化应用

容器本身可以通过`cggroups`进行资源隔离，比如可以限制容器 CPU 和 MEM 的最大使用占比

比如下面的命令: 

- `--cpus 0.1` ：只能使用 10%的CPU
- `-m 512M`: 只能使用 512M 的内存

```shell
$ docker run --name tomcat --cpus 0.1 -m 512M -p 8080:8080 -itd feisky/tomcat:8
```

而一旦限制了资源，一些应用在容器中资源耗尽，就会挂掉或者工作缓慢，这时可以先查看docker 的日志：

```shell
$ docker logs -f tomcat
```

再查看 docker 中进程的状态

```shell
$ docker ps -a
# 或查看指定应用状态的 json 格式的输出
$ docker inspect tomcat -f '{{json .State}}' | jq
```

如果进程挂掉了，可以通过`dmesg`命令查看系统日志

```shell
# -H 输出格式更容易让人理解，比如时间戳会转换成可读格式
$ dmesg -H
```

可以不直接进入容器运行指令检查应用

```shell
# 查看堆内存，注意单位是字节
$ docker exec tomcat java -XX:+PrintFlagsFinal -version | grep HeapSize
# 查看内存
$ docker exec tomcat free -m
```

可以直接进入容器的 bash 终端，再运行别的指令检查应用

```shell
$ docker exec -it tomcat bash
```

也可以在启动容器时指定初始命令

```shell
$ docker run --name tomcat --cpus 0.1 -m 512M -e JAVA_OPTS='-Xmx512m -Xms512m' -p 8080:8080 -itd feisky/tomcat:8
```

同时使用`top`指令观察CPU 使用率

```shell
$ top
```

当发现异常进程后，再使用`pidstat`分析指定进程

```shell
# -t表示显示线程，-p指定进程号
$ pidstat -t -p 29457 1
```

## 丢包问题检查

丢包可能的地方很多，需要通过网络收发流程图来逐步排查

![img](/../../../../../../../media/2020-07-11-linux-perf-optimization-learning-7th-week-summary/dd5b4050d555b1c23362456e357dfffd.png)

- 在两台 VM 连接之间，可能会发生传输失败的错误，比如网络拥塞、线路错误等；
- 在网卡收包后，环形缓冲区可能会因为溢出而丢包；
- 在链路层，可能会因为网络帧校验失败、QoS 等而丢包；可使用`netstat -i`查看丢包统计
- 在 IP 层，可能会因为路由失败、组包大小超过 MTU 等而丢包； 可使用`netstat -i`查看mtu
- 在传输层，可能会因为端口未监听、资源占用超过内核限制等而丢包；可使用`netstat -s`查看协议统计
- 在套接字层，可能会因为套接字缓冲区溢出而丢包；
- 在应用层，可能会因为应用程序异常而丢包； 可使用`tcpdump -i eth0 -nn port 80`查看指定端口网络包
- 此外，如果配置了 iptables 规则，这些网络包也可能因为 iptables 过滤规则而丢包。可使用`iptables -t filter -nvL`查看过滤规则

## 内核线程CPU 利用率高

- 使用 `top`命令查看占用 CPU 的进程 id 
- 使用`perf record`采样，`$ perf record -a -g -p 9 -- sleep 30`,然后执行`perf report`展开调用堆栈查看

还可以使用[火焰图](http://www.brendangregg.com/flamegraphs.html)，更直观的展示

![img](/../../../../../../../media/2020-07-11-linux-perf-optimization-learning-7th-week-summary/68b80d299b23b0cee518001f78960f61.png)

- 横轴表示采样数和采样比例。横轴越宽，代表执行时间越长
- 纵轴表示调用栈。由下往上根据调用关系逐个展开。纵轴越高，代表调用栈越深
- 颜色深浅没有特殊含义，只为函数间的区分

上图表示 CPU 的繁忙情况，还可以查看内存的分配和释放情况。

生成火焰图的方法：

- 安装`FlameGraph`, `git clone https://github.com/brendangregg/FlameGraph`
- 一行代码生成火焰图：`$ perf script -i /root/perf.data | ./stackcollapse-perf.pl --all |  ./flamegraph.pl > ksoftirqd.svg`

## 动态追踪

通过探针机制，来采集内核或者应用程序的运行信息，从而可以不用修改内核和应用程序的代码，就获得丰富的信息，帮你分析、定位想要排查的问题。

真实的线上项目，不能下断点，也不能停机进行分析，使用动态追踪技术可以在不停机，不修改代码的情况下直接进行分析，并且性能损耗相对于`ptrace`这种进程级跟踪方法会小很多（通常在 5% 或者更少）。

### ftrace

```shell
sudo apt-get install trace-cmd
```

以`ls`命令为例

```shell
sudo trace-cmd record -p function_graph -g do_sys_open -O funcgraph-proc ls
sudo trace-cmd report
```

## perf

安装调试符号信息

```shell
sudo apt install ubuntu-dbgsym-keyring
```

