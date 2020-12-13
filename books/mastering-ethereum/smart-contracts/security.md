---
title: 智能合约的安全漏洞
permalink: /books/masterring-ethereum/smart-contract/security
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

原因：

攻击者调用漏洞合约让其发送 eth 到攻击者合约，会触发攻击者合约的`payable`类型函数，如果攻击者合约在`payable`函数继续回调漏洞合约进行发送 eth，
漏洞合约可能会陷入无限递归，并且会不断发送 eth 到攻击者合约。

解决方法：

1. 使用 transfer 方法发送 eth，因为 transfer 方法只会使用 2300gas，不足以运行更复杂的递归调用
2. 确保在发送 eth 前就将状态修改掉，将发送 eth 作为最后一条语句
3. 使用 mutex 阻止重入

真实案件：

DAO 重入攻击，导致 ETH 分叉，成为了 ETC 和 ETH 两条链

## 溢出问题(Overflow)
