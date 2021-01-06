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

