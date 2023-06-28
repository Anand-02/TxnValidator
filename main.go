package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

var BlockCapacity = 3
var TxnNo = 0
var LastHash string
var ch = make(chan Txn, 10)
var tmpch = make(chan Txn)
var block = Block{
	BlockNo:       0,
	PrevBlockHash: LastHash,
}

func (newTxn Txn) DeriveHash(t details) {
	info, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err, t)
	}
	hash := sha256.Sum256((info))
	bs := fmt.Sprintf("%x", hash)
	newTxn.Data.Hash = bs
	tmpch <- newTxn
}

func (b Block) DeriveHash() {
	info, _ := json.Marshal(b)
	hash := sha256.Sum256([]byte(info))
	bs := fmt.Sprintf("%x", hash)
	LastHash = bs
}

func main() {
	SetKeys()
	InitChain()
	router.POST("/post", Handler)
	router.GET("/blocks", GetAllBlocks)
	router.GET("/blocks/:blockNumber", GetBlock)
	err := router.Run("localhost:8000")
	if err != nil {
		fmt.Println(err)
	}
}
