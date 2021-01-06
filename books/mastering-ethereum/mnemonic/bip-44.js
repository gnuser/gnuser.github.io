var bitcore = require('bitcore-lib');
var EthereumBip44 = require('ethereum-bip44');
// create a new master private key
var key = bitcore.HDPrivateKey();
// create the hd wallet
var wallet = new EthereumBip44(key);
// output the first address
console.log(wallet.getAddress(0));
// output the second address
console.log(wallet.getAddress(1));
