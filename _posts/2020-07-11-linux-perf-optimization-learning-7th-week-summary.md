---
key: linux-perf-optimization-learning-7th-week-summary
title: Linux性能学习第七周-综合实战
date: 2020-07-11 12:11:35 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
sidebar:
  nav: lpo-note
permalink: /lpo-note/7th-week
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

### perf

安装调试符号信息， 参考[https://wiki.ubuntu.com/Debug%20Symbol%20Packages](https://wiki.ubuntu.com/Debug Symbol Packages)

```shell
codename=$(lsb_release -c | awk  '{print $2}')
sudo tee /etc/apt/sources.list.d/ddebs.list << EOF
deb http://ddebs.ubuntu.com/ ${codename}      main restricted universe multiverse
deb http://ddebs.ubuntu.com/ ${codename}-security main restricted universe multiverse
deb http://ddebs.ubuntu.com/ ${codename}-updates  main restricted universe multiverse
deb http://ddebs.ubuntu.com/ ${codename}-proposed main restricted universe multiverse
EOF
sudo apt install ubuntu-dbgsym-keyring
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys F2EDC64DC5AEE1F6B9C621F0C8CAB6595FDFF622
sudo apt-get update
sudo apt-get install linux-image-$(uname -r)-dbgsym
```

以观察内核函数 `do_sys_open` 为例子

- 查看函数的定义，参数的类型

```shell
$ sudo perf probe -V do_sys_open
Available variables at do_sys_open
        @<do_sys_open+0>
                char*   filename
                int     dfd
                int     flags
                struct open_flags       op
                umode_t mode
```

这里我们比较关注的是`filename`，也就是希望看到打开的文件名称，如果查看不到以上内容，说明调试符号没有安装好。

- 添加探针

```shell
sudo perf probe --add 'do_sys_open filename:string'
```

注意，如果只写函数名称，是看不到函数打开的文件名的

- 采样

```shell
sudo perf record -e probe:do_sys_open -aR ls
```

- 查看结果

```shell
$ sudo perf script
            perf 23321 [002] 1216304.406239: probe:do_sys_open: (ffffffff93276e90) filename_string="/proc/23324/status"
              ls 23324 [001] 1216304.406588: probe:do_sys_open: (ffffffff93276e90) filename_string="/etc/ld.so.cache"
              ls 23324 [001] 1216304.406603: probe:do_sys_open: (ffffffff93276e90) filename_string="/lib/x86_64-linux-gnu/libselinux.so.1"
...
              ls 23324 [001] 1216304.407014: probe:do_sys_open: (ffffffff93276e90) filename_string="/usr/lib/locale/zh.UTF-8/LC_CTYPE"
              ls 23324 [001] 1216304.407017: probe:do_sys_open: (ffffffff93276e90) filename_string="/usr/lib/locale/zh.utf8/LC_CTYPE"
              ls 23324 [001] 1216304.407021: probe:do_sys_open: (ffffffff93276e90) filename_string="/usr/lib/locale/zh/LC_CTYPE"
              ls 23324 [001] 1216304.407034: probe:do_sys_open: (ffffffff93276e90) filename_string="."
```

- 删除探针

```shell
sudo perf probe --del probe:do_sys_open
```

### eBPF 和 BCC

eBPF 就是 Linux 版的 DTrace，可以通过 C 语言自由扩展

BCC（BPF Compiler Collection,把这些过程通过 Python 抽象起来,使用 BCC 进行动态追踪时，编写简单的脚本就可以了。

- 安装 BCC

```shell

# Ubuntu
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 4052245BD4284CDD
echo "deb https://repo.iovisor.org/apt/$(lsb_release -cs) $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/iovisor.list
sudo apt-get update
sudo apt-get install bcc-tools libbcc-examples linux-headers-$(uname -r)

```

- 使用 python 脚本监控`do_sys_open`

```python
from bcc import BPF
# define BPF program (""" is used for multi-line string).
# '#' indicates comments for python, while '//' indicates comments for C.
prog = """
#include <uapi/linux/ptrace.h>
#include <uapi/linux/limits.h>
#include <linux/sched.h>
// define output data structure in C
struct data_t {
    u32 pid;
    u64 ts;
    char comm[TASK_COMM_LEN];
    char fname[NAME_MAX];
};
BPF_PERF_OUTPUT(events);

// define the handler for do_sys_open.
// ctx is required, while other params depends on traced function.
int hello(struct pt_regs *ctx, int dfd, const char __user *filename, int flags){
    struct data_t data = {};
    data.pid = bpf_get_current_pid_tgid();
    data.ts = bpf_ktime_get_ns();
    if (bpf_get_current_comm(&data.comm, sizeof(data.comm)) == 0) {
        bpf_probe_read(&data.fname, sizeof(data.fname), (void *)filename);
    }
    events.perf_submit(ctx, &data, sizeof(data));
    return 0;
}
"""
# load BPF program
b = BPF(text=prog)
# attach the kprobe for do_sys_open, and set handler to hello
b.attach_kprobe(event="do_sys_open", fn_name="hello")

# process event
start = 0
def print_event(cpu, data, size):
    global start
    event = b["events"].event(data)
    if start == 0:
            start = event.ts
    time_s = (float(event.ts - start)) / 1000000000
    print("%-18.9f %-16s %-6d %-16s" % (time_s, event.comm, event.pid, event.fname))

# loop with callback to print_event
b["events"].open_perf_buffer(print_event)


# print header
print("%-18s %-16s %-6s %-16s" % ("TIME(s)", "COMM", "PID", "FILE"))
# start the event polling loop
while 1:
    try:
        b.perf_buffer_poll()
    except KeyboardInterrupt:
        exit()
```

- 开始监控

```shell
sudo python trace-open.py
```

### SystemTap 和 sysdig

从稳定性上来说，SystemTap 只在 RHEL 系统中好用，在其他系统中则容易出现各种异常问题

sysdig 则是随着容器技术的普及而诞生的，主要用于容器的动态追踪。sysdig 汇集了一些列性能工具的优势，可以说是集百家之所长。 sysdig 的特点： sysdig = strace + tcpdump + htop + iftop + lsof + docker inspect。

## 几个常见的动态追踪使用场景

![img](/../../../../../../../media/2020-07-11-linux-perf-optimization-learning-7th-week-summary/5a2b2550547d5eaee850bfb806f76625.png)



eBPF编写可以参考[https://docs.cilium.io/en/stable/bpf/](https://docs.cilium.io/en/stable/bpf/)



