package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/btcec"
)

func main() {
	dataHash := sha256.Sum256([]byte("ethereum"))

	// 准备私钥
	pkeyb, err := hex.DecodeString("289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032")
	if err != nil {
		log.Fatalln(err)
	}
	// 基于secp256k1的私钥
	privk, _ := btcec.PrivKeyFromBytes(btcec.S256(), pkeyb)

	// 对内容的 hash 进行签名
	sigInfo, err := privk.Sign(dataHash[:])
	if err != nil {
		log.Fatal(err)
	}
	// 获得DER格式的签名
	sig := sigInfo.Serialize()
	fmt.Println("sig length:", len(sig))
	fmt.Println("sig hex:", hex.EncodeToString(sig))
}
