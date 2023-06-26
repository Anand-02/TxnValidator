package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

var BlockCapacity = 3
var TxnNo = 0
var LastHash [32]byte

// var db = openDB()
func (newTxn Txn) DeriveHash(t details) {
	info, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err , t)
	}
	//fmt.Println(info)
	hash := sha256.Sum256((info))
	
	newTxn.Data.Hash = hash
	fmt.Println(newTxn)
	ch <- newTxn
}

func (b Block) DeriveHash() {
	info, _ := json.Marshal(b)
	hash := sha256.Sum256([]byte(info))
	fmt.Printf("%T", hash)
	LastHash = hash
}

func main() {
	// defer openDB().Close()
	SetKeys()
	createRoute()
}
