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
    
