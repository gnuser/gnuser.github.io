```shell
# 观察 TCP 连接数
$ ss -s
Total: 177 (kernel 1565)
TCP:   1193 (estab 5, closed 1178, orphaned 0, synrecv 0, timewait 1178/0), ports 0

Transport Total     IP        IPv6
*    1565      -         -
RAW    1         0         1
UDP    2         2         0
TCP    15        12        3
INET    18        14        4
FRAG    0         0         0
```



```shell
# 查看套接字的队列大小
$ ss -ltnp
State     Recv-Q     Send-Q            Local Address:Port            Peer Address:Port
LISTEN    10         10                      0.0.0.0:80                   0.0.0.0:*         users:(("nginx",pid=10491,fd=6),("nginx",pid=10490,fd=6),("nginx",pid=10487,fd=6))
LISTEN    7          10                            *:9000                       *:*         users:(("php-fpm",pid=11084,fd=9),...,("php-fpm",pid=10529,fd=7))
```

