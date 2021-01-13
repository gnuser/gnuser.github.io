---
title: transactions
permalink: /books/mastering-ethereum/transactions
key: transactions
layout: article
sidebar:
  nav: me
aside:
  toc: true
---

交易

<!--more-->

## Transactions

区块链上实际存放的都是 transactions，通过 transaction 的内容可以更新账户或者合约里的状态，transaction 的执行是有顺序的，原子化的，和 eventsourcing 里面的 event 序列类似。

## 数据结构

**Nonce**
保证交易顺序执行，防止重放攻击

**Gas price**
单位为 wei，决定交易打包快慢，可以参考选取 gas 费用[https://www.gasnow.org/](https://www.gasnow.org/)

**Gas limit**
最大允许 gas 使用量，可以使用 estimateGas 预估

**Recipient**
交易目的地址

**Value**
转账 eth 数量，可以为 0(一般为合约调用)

**Data**
附加的动态二进制数据

**v,r,s**
椭圆曲线签名信息，可以参考前面的章节

## 什么是 RLP？

Recursive Length Prefix

transaction序列化格式

## transaction 数据并没有包含from字段

因为我们可以从v,r,s推出public key，所以也可以得到from地址

同样，block number和transaction id也没有包含在transaction数据里

## Nonce值的获取方式

```
web3.eth.getTransactionCount("0x9e713963a92c02317a681b9bb3065a8249de124f")
40
```

- nonce从0开始计数
- 如果连续发出交易，要注意`getTransactionCount`获取到的值可能不对，解决方案是换成`Parity`节点的`parity_nextNonce`，或者自己维护好`nonce` 值

- 如果nonce没有按顺序发出，比如当前是0，如果发出了3的transaction，3的会存在节点的mempool里，直到漏掉的transaction全部上链

## transaction没有并发上链一说

多个transaction可以同时广播，但不会同时执行，顺序一般由gas fee决定（transaction在区块内的index）

同时同一地址的transaction顺序由nonce值保证，所以当你同时发出多笔交易时，如果后面的交易依赖于前面的交易结果（比如approve），也不需要担心顺序的问题

## 如果transaction一直pending怎么办？

使用同样的nonce，发送另一笔transaction覆盖之前的transaction，所幸的是你之前的transaction并不会额外消耗gas fee

## gas price获取方式

1. 调用rpc接口

```
> web3.eth.getGasPrice(console.log)
> null BigNumber { s: 1, e: 10, c: [ 10000000000 ] }
```

2. 调用第三方服务api

```
https://www.gasnow.org/api/v3/gas/data?utm_source=web
```

## gas limit预估

- 普通的eth转账： 21000
- 合约交易：使用estimateGas可以预估，但不能保证一定足够，但是如果没有使用完的，会退还给账号

## Recipient如果不是20字节会怎么样？

## transaction的value和data



