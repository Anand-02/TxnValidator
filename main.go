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
var block = Block{
	BlockNo:       0,
	PrevBlockHash: LastHash,
}

// var db = openDB()
func (newTxn Txn) DeriveHash(t details) {
	info, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err, t)
	}
	//fmt.Println(info)
	hash := sha256.Sum256((info))
	bs := fmt.Sprintf("%x", hash)

	newTxn.Data.Hash = bs
	// fmt.Println(newTxn)
	ch <- newTxn
}

func (b Block) DeriveHash() {
	info, _ := json.Marshal(b)
	hash := sha256.Sum256([]byte(info))
	bs := fmt.Sprintf("%x", hash)
	// fmt.Printf("%T", hash)
	LastHash = bs
	fmt.Println(LastHash)
}

func main() {
	// defer openDB().Close()
	SetKeys()
	go InitChain()
	router.POST("/post", Handler)
	err := router.Run("localhost:8000")
	Err(err)
}
