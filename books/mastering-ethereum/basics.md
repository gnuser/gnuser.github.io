---
title: Ethereum Basics
permalink: /books/mastering-ethereum/basics
key: basics
layout: article
sidebar:
  nav: me
aside:
  toc: true
---

如何使用钱包，创建交易，部署简单的智能合约

<!--more-->

## 以太坊的货币基础单位

> Ethereum is the system, ether is the currency.

1 ETH = 10**18 wei

## 钱包类型

主要区别就是支持的币种数量，以及私钥是否保存本地

- 浏览器插件 MetaMask

[https://metamask.io](https://metamask.io)

目前最流行的钱包，chrome和firefox都有相应的插件，各个dapp首选的钱包，但只支持eth

- App Imtoken

国内最流行的手机端钱包，支持冷钱包扫码方式，支持硬件钱包，有内嵌dapp平台，支持btc，eth，eos，tether

- 网页版钱包 Blockchain.com

私钥由平台管理，支持btc，eth等

## 钱包安全

- 私钥或者助记词妥善保管，不可泄漏
- 只要有助记词，密钥丢失也可重置
- 大额资产最好使用冷钱包保管
- 转账时最好先小额支付一笔，确定无误后再转出

## 可选的网络模式

- 主网络

真正的eth资产

- ropsten
- kovan
- rinkeby
- goerli

都是测试网络，eth可以通过faucet网络获取，但最容易获得的是ropsten [https://faucet.metamask.io/](https://faucet.metamask.io/)

- 私有节点网络

自建节点，可以搭建主网节点，也可以搭建测试网络，或者私有网络节点

## 账户模型

以太坊包含两种账户，一种是存放eth的，我们称为EOA(externally owned accounts)，另一种是存放token的，我们称为合约账户。

EOA有私钥，合约账户没有

EOA和合约账户都能接收eth和token

## 合约初试

[https://github.com/gnuser/simple-dapp-tutorial](https://github.com/gnuser/simple-dapp-tutorial)






