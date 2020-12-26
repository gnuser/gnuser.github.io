require 'openssl'
require 'base16'
require 'digest/sha3'
def eth_address(public_key)
  s = public_key[2, 128]
  s.downcase!
  puts s
  s = Base16.decode16(s)
  puts s
  h = Digest::SHA3.hexdigest(s, 256)
  puts h
  a = '0x' + h[-40..-1]
  return a
end
ec = OpenSSL::PKey::EC.new('secp256k1')
ec.generate_key
public_key = ec.public_key.to_bn.to_s(16)
private_key = ec.private_key.to_s(16)
eth_address = eth_address(public_key)
puts "address: #{eth_address}"
puts "private_key: #{private_key}"
puts "public_key: #{public_key}"

puts eth_address(
       '0x6e145ccef1033dea239875dd00dfb4fee6e3348b84985c92f103444683bae07b83b5c38e5e2b0c8529d7fa3f64d46daa1ece2d9ac14cab9477d042c84c32ccd0'
     )
