const bip39 = require('bip39')
 
// defaults to BIP39 English word list
// uses HEX strings for entropy
const mnemonic = bip39.entropyToMnemonic('b928973a6742a1605ffc7d3994dbb7bec9544226f21a5aaab212da4d15fe2e76')

console.log(mnemonic)

seed = bip39.mnemonicToSeedSync('rice dwarf soldier soldier claw rabbit lend moral define plug unknown laugh next canal orange drive follow few lucky region spend yellow right source', 'gnuser')

console.log('seed:', seed.toString('hex'));
