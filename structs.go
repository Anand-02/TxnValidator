package main

import "time"

type Block struct {
	BlockNo       int
	PrevBlockHash string
	Txns          []Txn
	TimeStamp     time.Time
	CommitStatus  bool
}

type Txn struct {
	Id   string
	Data Data
}

type Data struct {
	TxnID   int
	Value   int     `json:"val"`
	Version float32 `json:"ver"`
	Valid   bool
	Hash    string
}

type details struct {
	Value   int
	Version float32
}
