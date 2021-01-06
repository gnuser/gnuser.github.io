require 'bip_mnemonic'

entropy =
  BipMnemonic.to_entropy(
    mnemonic:
      'rice dwarf soldier soldier claw rabbit lend moral define plug unknown laugh next canal orange drive follow few lucky region spend yellow right source',
    language: 'english'
  )

p entropy
