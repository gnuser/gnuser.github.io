---
title: wallets
permalink: /books/mastering-ethereum/wallets
key: wallets
layout: article
sidebar:
  nav: me
aside:
  toc: true
---

以太坊钱包

<!--more-->

## 关于钱包

这里我们说的钱包都是非托管钱包，也就是私钥在本地存放的钱包。
我们的资产记录记录在以太坊链上，当使用时，我们使用私钥对交易签名，并广播上链。
和传统的银行账号不同，所有的交易都是由私钥拥有者授权进行，而所有的账户的资产以及交易记录都是公开透明的。
当你把私钥导入其他钱包时，通常情况就可以在其他钱包查看资产以及签名交易。

## 钱包类型

- nondeterministic wallet
  每个私钥互不相关
- deterministic wallet
  所有的私钥都从一个主私钥衍生的

## keystore 文件

文件格式如下:

```json
{
  "address": "001d3f1ef827552ae1114027bd3ecf1f086ba0f9",
  "crypto": {
    "cipher": "aes-128-ctr",
    "ciphertext": "233a9f4d236ed0c13394b504b6da5df02587c8bf1ad8946f6f2b58f055507ece",
    "cipherparams": {
      "iv": "d10c6ec5bae81b6cb9144de81037fa15"
    },
    "kdf": "scrypt",
    "kdfparams": {
      "dklen": 32,
      "n": 262144, # hash计算次数
      "p": 1,
      "r": 8,
      "salt": "99d37a47c7c9429c66976f643f386a61b78b97f3246adca89abe4245d2788407"
    },
    "mac": "594c8df1c8ee0ded8255a50caf07e8c12061fd859f4b7c76ab704b17c957e842"
  },
  "id": "4fcb2ba4-ccdb-424f-89d5-26cce304bf9c",
  "version": 3
}
```

## 助记词（BIP-39）

助记词就和私钥一样，只要拥有，就拥有对应地址的资产

生成主记词流程

![Generating entropy and encoding as mnemonic words](../media/wallets/bip39-part1.png)

1. 生成随机 128 位的 0，1 串 S
2. sha256(S)，截取长度为( S 的长度 / 32)作为checksum
3. 拼接S和checksum得到一个新串，命名为SC
4. 将SC切成12个11bit的数组
5. 每个11bit对应的值映射到单词字典上
6. 获得12个单词的助记词

**根据第一步随机位数不同，checksum的长度不同，助记词的长度也不同，可以看到如果是256位，对应的就是24个单词**

| Entropy (bits) | Checksum (bits) | Entropy **+** checksum (bits) | Mnemonic length (words) |
| -------------- | --------------- | ----------------------------- | ----------------------- |
| 128            | 4               | 132                           | 12                      |
| 160            | 5               | 165                           | 15                      |
| 192            | 6               | 198                           | 18                      |
| 224            | 7               | 231                           | 21                      |
| 256            | 8               | 264                           | 24                      |

接下来再把助记词转换为seed

7. 将助记词和salt拼接为PBKDF2
8. salt是`mnemonic + passphrase`,passphrase是用户提供的密码，可以为空
9. 对PBKDF2进行2048次`HMAC-SHA512`hash计算，生成一个512bit的seed

![From mnemonic to seed](../media/wallets/bip39-part2.png)

## 使用不同语言生成助记词

[python-mnemonic](https://github.com/trezor/python-mnemonic)

```
pip install mnemonic
```



[ConsenSys/eth-lightwallet](https://github.com/ConsenSys/eth-lightwallet)

[npm/bip39](https://www.npmjs.com/package/bip39)

https://iancoleman.io/bip39/

