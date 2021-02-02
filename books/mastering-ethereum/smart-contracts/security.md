---
title: 智能合约的安全漏洞
permalink: /books/mastering-ethereum/smart-contract/security
key: smart-contract/security
layout: article
sidebar:
  nav: me
aside:
  toc: true
---

智能合约的安全漏洞

<!--more-->

## 重入问题(Reentrancy)

### 原因：

攻击者调用漏洞合约让其发送 eth 到攻击者合约，会触发攻击者合约的`payable`类型函数，如果攻击者合约在`payable`函数继续回调漏洞合约进行发送 eth，
漏洞合约可能会陷入无限递归，并且会不断发送 eth 到攻击者合约。

### 解决方案：

1. 使用 transfer 方法发送 eth，因为 transfer 方法只会使用 2300gas，不足以运行更复杂的递归调用
2. 确保在发送 eth 前就将状态修改掉，将发送 eth 作为最后一条语句
3. 使用 mutex 阻止重入

### 真实案例：

DAO 重入攻击，导致 ETH 分叉，成为了 ETC 和 ETH 两条链

## 溢出问题(Overflow)

## Unexpected Ether

### 原因:

攻击者通过self-destruct和pre-sent方式向合约发送eth，而不需要调用payable函数

### 解决方案:

## Default Visibilities

- 对于其他合约，方法的可见性由`internally`和`externally`控制
- 对于EOA用户，方法的可见性由`public`和`private`控制

### 原因：

而默认的方法可见性是`public`

### 解决方案:

每一个方法都显式的指明可见性，不要用默认的方式

### 真实案例:

#### Parity Multisig Wallet (First Hack), 损失$31M

## Entropy Illusion

### 原因：

一些博彩类游戏需要生成随机数，而链上随机数是不安全的，如果使用未来区块中包含的比如hashes,timestaps,block numbers, gas limits.这些数据是可以被矿工操控的，矿工可以选择不打包上链不利于他的区块。如果使用过去的或者当前的区块数据产生随机数，攻击者可以利用同一区块生成的随机数相同，将收益翻倍。

### 解决方案：

随机数的熵必须来自链下的数据源，可以通过oracle，不要使用区块数据作为熵

### 真实案例：

PRNG合约(pseudorandom number generator)

## External Contract Referencing

### 原因：

合约重用时，权限控制失败，被攻击者替换了调用合约地址

### 解决方案：

- 使用`new`创建库合约
- hardcode库合约地址
- 使用`time-lock`或者`voting`机制

### 真实案例：Renentrancy Honey Pot

## Short Address/Parameter Attack

### 原因：

没有检查地址的长度，被攻击者利用

### 解决方案：

检查所有的合约参数调用

## Unchecked CALL Return Values

### 原因：

有三种方式可以发送eth给其他合约

- transfer
- send
- call

`call`和`send`会返回bool说明调用是否成功，如果调用失败，也不会回滚，只是返回`false`

### 解决方案：

总是使用`transfer`，而不是使用`send`,因为前者会回滚，否则需要检查`send`的返回值

### 真实案例: Etherpot and King of the Ether

## Race Conditions/Front Running

### 原因：

攻击者可以监控mempool里的transactions，并使用更高的gasPrice进行抢跑获取利益。

矿工甚至可以选择性打包来控制区块数据。

### 解决方案：

- 设置`gasPrice`上限
- 使用`commit-reveal`方案
- 使用`submarine sends`

### 真实案例：ERC20 and Banchor

- 两次approve操作中间被抢跑，相当于两次approve都有效
- Banchor和uniswap都有这样的问题，抢跑可以拿到更好的价格，[抢跑python实现](https://hackernoon.com/front-running-bancor-in-150-lines-of-python-with-ethereum-api-d5e2bfd0d798)

## Denial Of Service













