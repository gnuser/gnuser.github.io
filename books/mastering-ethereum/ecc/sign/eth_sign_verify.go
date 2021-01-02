package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	decodeHex := func(s string) []byte {
		b, err := hex.DecodeString(s)
		if err != nil {
			log.Fatal(err)
		}
		return b
	}
	dataHash := sha256.Sum256([]byte("ethereum"))
	sig := decodeHex(
		"7912f50819764de81ab7791ab3d62f8dabe84c2fdb2f17d76465d28f8a968f7355fbb6cd8dfc7545b6258d4b032753b2074232b07f3911822b37f024cd10116600")
	pubkey := decodeHex(
		"037db227d7094ce215c3a0f57e1bcc732551fe351f94249471934567e0f5dc1bf7")

	ok := crypto.VerifySignature(pubkey, dataHash[:], sig[:len(sig)-1])
	fmt.Println("verify pass?", ok)
}
