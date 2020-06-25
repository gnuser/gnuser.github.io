---
title: ab-HTTP性能
permalink: /lpo/network/ab
key: network/ab
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

要测试 HTTP 的性能，也有大量的工具可以使用，比如 ab、webbench 等，都是常用的 HTTP 压力测试工具。其中，ab 是 Apache 自带的 HTTP 压测工具，主要测试 HTTP 服务的每秒请求数、请求延迟、吞吐量以及请求延迟的分布情况等。

<!--more-->

## 安装
```shell
# Ubuntu
$ sudo apt-get install -y apache2-utils
# CentOS
$ yum install -y httpd-tools
```

## 在目标机器启动`Nginx`服务
```shell
$ docker run -p 80:80 -itd nginx
```

## 在另一台机器使用`ab`测试
```shell
# -c表示并发请求数为1000，-n表示总的请求数为10000
$ ab -c 1000 -n 10000 http://172.18.9.134/
This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 172.18.9.134 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        nginx/1.14.0
Server Hostname:        172.18.9.134
Server Port:            80

Document Path:          /
Document Length:        612 bytes

Concurrency Level:      1000
Time taken for tests:   0.510 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      8540000 bytes
HTML transferred:       6120000 bytes
Requests per second:    19617.61 [#/sec] (mean)
Time per request:       50.975 [ms] (mean)
Time per request:       0.051 [ms] (mean, across all concurrent requests)
Transfer rate:          16360.78 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        3   10   4.8     10      24
Processing:     2   21  28.8     14     222
Waiting:        2   17  28.9     11     222
Total:          8   31  30.5     25     242

Percentage of the requests served within a certain time (ms)
  50%     25
  66%     32
  75%     37
  80%     39
  90%     43
  95%     62
  98%    103
  99%    228
 100%    242 (longest request)
```
ab 的测试结果分为三个部分，分别是请求汇总、连接时间汇总还有请求延迟汇总

- Requests per second 为 `19617.61`
- 每个请求的延迟（Time per request）分为两行，第一行的 50.975 ms 表示平均延迟，包括了线程运行的调度时间和网络请求响应时间，而下一行的 0.051ms ，则表示实际请求的响应时间；
- Transfer rate 表示吞吐量（BPS）为 16360.78 KB/s。

连接时间汇总部分，则是分别展示了建立连接、请求、等待以及汇总等的各类时间，包括最小、最大、平均以及中值处理时间。

最后的请求延迟汇总部分，则给出了不同时间段内处理请求的百分比，比如， 90% 的请求，都可以在 43ms 内完成。