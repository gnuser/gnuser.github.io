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

```python
from mnemonic import Mnemonic
from binascii import hexlify, unhexlify

mnemo = Mnemonic("english")

passphrase = "gnuser"
words = mnemo.generate(strength=256)
seed = mnemo.to_seed(words, passphrase=passphrase)
entropy = mnemo.to_entropy(words)

print("mnemonic:", words)
print("seed:", hexlify(seed))
print("entropy:", hexlify(entropy))
```

```
('mnemonic:', u'rice dwarf soldier soldier claw rabbit lend moral define plug unknown laugh next canal orange drive follow few lucky region spend yellow right source')
('seed:', '1fe60435217a0ead9384da5abbfabc32c9951d0b86ae127920db2fe3a6950996d161413e5bcb7eaec5b2975b07ba8e8a5e5912ecb1ebadd703fee78bd19ffd80')
('entropy:', 'b928973a6742a1605ffc7d3994dbb7bec9544226f21a5aaab212da4d15fe2e76')
```

[ConsenSys/eth-lightwallet](https://github.com/ConsenSys/eth-lightwallet)

[npm/bip39](https://www.npmjs.com/package/bip39)

```shell
npm install bip39
```

```javascript
const bip39 = require('bip39')

const mnemonic = bip39.entropyToMnemonic('b928973a6742a1605ffc7d3994dbb7bec9544226f21a5aaab212da4d15fe2e76')

console.log(mnemonic)

seed = bip39.mnemonicToSeedSync('rice dwarf soldier soldier claw rabbit lend moral define plug unknown laugh next canal orange drive follow few lucky region spend yellow right source', 'gnuser')

console.log('seed:', seed.toString('hex'));
```

[https://github.com/sreekanthgs/bip_mnemonic](https://github.com/sreekanthgs/bip_mnemonic)

```ruby
require 'bip_mnemonic'

entropy =
  BipMnemonic.to_entropy(
    mnemonic:
      'rice dwarf soldier soldier claw rabbit lend moral define plug unknown laugh next canal orange drive follow few lucky region spend yellow right source',
    language: 'english'
  )

p entropy
```

https://iancoleman.io/bip39/

## 创建HD钱包

HD钱包的好处

- 可以不需要私钥生成很多的扩展公钥
- 对于像交易所这样的应用来说，就不需要生成并存放私钥在节点服务器上
- 签名的时候可以使用扩展私钥进行离线签名，更加安全

HD钱包的风险：

- 一旦扩展私钥被盗，所有的子私钥都意味着被盗

### HD钱包的path举例

m代表私钥路径，M代表公钥路径

| HD path     | Key described                                                |
| ----------- | ------------------------------------------------------------ |
| m/0         | The first (0) child private key of the master private key (m) |
| m/0/0       | The first grandchild private key of the first child (m/0)    |
| m/0'/0      | The first normal grandchild of the first *hardened* child (m/0') |
| m/1/0       | The first grandchild private key of the second child (m/1)   |
| M/23/17/0/0 | The first great-great-grandchild public key of the first great-grandchild of the 18th grandchild of the 24th child |

### BIP-44简化path

```
m / purpose' / coin_type' / account' / change / address_index
```

purpose通常使用`44'`

coin_type: 参考[SLIP0044](https://github.com/satoshilabs/slips/blob/master/slip-0044.md)，eth使用`60`,etc使用`61`,btc使用`0`

account: 用户自定义，可以按组织划分

change: 可以用`0`代表收款地址，`1`代表零钱地址，ethereum只使用`0`

### Hardened child key

`0`和`0'`的区别，后者代表是`hardened child`,从`2^31`后开始计数，所以`0'`代表的是`2^31+0`

### BIP-44举例

| HD path                         | Key described                                                |
| ------------------------------- | ------------------------------------------------------------ |
| M/44&#x27;/60&#x27;/0&#x27;/0/2 | The third receiving public key for the primary Ethereum account |
| M/44&#x27;/0&#x27;/3&#x27;/1/14 | The 15th change-address public key for the 4th Bitcoin account |
| m/44&#x27;/2&#x27;/0&#x27;/0/1  | The second private key in the Litecoin main account, for signing transactions |

## BIP-44钱包生成

```shell
pip install rlp eth_utils two1 
pip install pycrypto 
pip install pycryptodome
# 确保pycrypto先于pycryptodome安装 https://github.com/pycrypto/pycrypto/issues/297
```

下载https://github.com/michailbrynard/ethereum-bip44-python/blob/master/crypto.py

```python
from crypto import HDPrivateKey, HDPublicKey, HDKey

master_key, mnemonic = HDPrivateKey.master_key_from_entropy()
print('BIP32 Wallet Generated.')
print('Mnemonic Secret: ' + mnemonic)

root_keys = HDKey.from_path(master_key,"m/44'/60'/0'")
acct_priv_key = root_keys[-1]

acct_pub_key = acct_priv_key.public_key
print('Account Master Public Key (Hex): ' + acct_pub_key.to_hex())
print('XPUB format: ' + acct_pub_key.to_b58check())

keys = HDKey.from_path(acct_pub_key,'{change}/{index}'.format(change=0, index=0))
address = keys[-1].address()
print('Account address: ' + address)

for i in range(10):
    keys = HDKey.from_path(acct_priv_key,'{change}/{index}'.format(change=0, index=i))
    private_key = keys[-1]
    public_key = private_key.public_key
    print("Index %s:" % i)
    print("  Private key (hex, compressed): " + private_key._key.to_hex())
    print("  Address: " + private_key.public_key.address())
```

```shell
$ python bip-44.py
BIP32 Wallet Generated.
Mnemonic Secret: primary average curve concert reward bench song digital auction unknown local demand
Account Master Public Key (Hex): 0488b21e034825c426800000004a57f9542f47e336a09e8e3e5baa910deceef9c142f7134517c99041f43c9d3002e4f80c9e8ee46cd6ad5f4ef39df9891fc3fc2f0f911a1e432c4130454ed97783
XPUB format: xpub6CBXU6YW4wfwKPQP7BoSJ659HonPpCQXzxAujyWaWberJF1ZdTp3ZUDUeyybXruSuCb2jqCgnNP1qUZ21T3D6ges3PPJZe8SwaxJsZ1HNma
Account address: 0x0585af690792e348d9609c2632533c1a1703b448
Index 0:
  Private key (hex, compressed): 6246a02fa684afea7c197e8c6792aa9d1f72a618897ed7c4e09d436146548949
  Address: 0x0585af690792e348d9609c2632533c1a1703b448
  Private key (hex): 6246a02fa684afea7c197e8c6792aa9d1f72a618897ed7c4e09d436146548949
Index 1:
  Private key (hex, compressed): 26777d44ff7e22a9fcb749324f32c91a11557a21b18e577220f6aefa9cc19f43
  Address: 0x71f196c58ddc33d2cccc25c30c437318808db525
  Private key (hex): 26777d44ff7e22a9fcb749324f32c91a11557a21b18e577220f6aefa9cc19f43
Index 2:
  Private key (hex, compressed): 51395cdfbfb765019630e9960a9e04a0a23639ba2c702831ea937b9c6834583b
  Address: 0x3f1805c5812696e784d33ade67b5ac9ebe96bdbf
  Private key (hex): 51395cdfbfb765019630e9960a9e04a0a23639ba2c702831ea937b9c6834583b
Index 3:
  Private key (hex, compressed): 1dbd093dffbf3051f633ea39a795a714fd895236367f42053238fc43d1a3426e
  Address: 0xbb96e630ca83fedd0eed539434e9e13c70a87f26
  Private key (hex): 1dbd093dffbf3051f633ea39a795a714fd895236367f42053238fc43d1a3426e
Index 4:
  Private key (hex, compressed): 691eb07fbfe652edb19dc3ebefbe156840c9cb848243a1d7f93ffc959afae5f1
  Address: 0xe50d48a55db9ea1ec7602e0253c79df006db203d
  Private key (hex): 691eb07fbfe652edb19dc3ebefbe156840c9cb848243a1d7f93ffc959afae5f1
Index 5:
  Private key (hex, compressed): 7418edc8982f727e69b3dc68aaba2c45d2afc5e28e3e27ef44e31f1419595474
  Address: 0x45d8b1a75c866ef9fbe0517c1649a370c5662194
  Private key (hex): 7418edc8982f727e69b3dc68aaba2c45d2afc5e28e3e27ef44e31f1419595474
Index 6:
  Private key (hex, compressed): 493e2cf8b2a3c6ded91aec875369d5a48071dad462b460ff6a9c77b7092e56c1
  Address: 0x81a0e7af8b90c1b2a61deb65a305f6632c4f9962
  Private key (hex): 493e2cf8b2a3c6ded91aec875369d5a48071dad462b460ff6a9c77b7092e56c1
Index 7:
  Private key (hex, compressed): 6216613212f422f227641562740be6405dfc960a579b20c72bbde35244197f42
  Address: 0x926bd662a81c893951887ddd000699643a1b6a59
  Private key (hex): 6216613212f422f227641562740be6405dfc960a579b20c72bbde35244197f42
Index 8:
  Private key (hex, compressed): 651e06a527575a6ca62ea800ef7de18d0a5362da49dd2e9210b1989486639dfb
  Address: 0x4339db5c5069d74f7fe935e5c2069ec8a7d220db
  Private key (hex): 651e06a527575a6ca62ea800ef7de18d0a5362da49dd2e9210b1989486639dfb
Index 9:
  Private key (hex, compressed): d4972123972aae259db66e5c6ed2f86478430148fb4aa02ca4dcbadf0ae222ef
  Address: 0x6a70891dcc405b6c464885fba6d78092a6aef248
  Private key (hex): d4972123972aae259db66e5c6ed2f86478430148fb4aa02ca4dcbadf0ae222eface
```

