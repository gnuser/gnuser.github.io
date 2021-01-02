package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	dataHash := sha256.Sum256([]byte("ethereum"))

	// 准备私钥
	pkeyb, err := hex.DecodeString("289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032")
	if err != nil {
		log.Fatalln(err)
	}
	// 基于secp256k1的私钥
	pkey, err := crypto.ToECDSA(pkeyb)
	if err != nil {
		log.Fatalln(err)
	}
	// 签名
	sig, err := crypto.Sign(dataHash[:], pkey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("sig length:", len(sig))
	fmt.Println("sig hex:", hex.EncodeToString(sig))
}
