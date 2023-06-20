package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// type Block struct {
// 	Hash     []byte
// 	PrevHash []byte
// 	Ledger   ledger
// }

// type ledger struct {
// 	BlockNo int
// 	Txns    []txn
// }

// type txn struct {
// 	Id      string
// 	Data    details
// 	TxnHash []byte
// }

type details struct {
	Value   int
	Version float32
}

func (t *details) DeriveHash() [32]uint8 {
	info, _ := json.Marshal(t)
	hash := sha256.Sum256([]byte(info))
	fmt.Printf("%T", hash)
	return hash
}

func main() {
	SetKeys()
}
