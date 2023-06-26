package main

import "time"

type Block struct {
	BlockNo       int
	PrevBlockHash [32]byte
	Txns          []Txn
	TimeStamp     time.Time
	CommitStatus  bool
}

type Txn struct {
	Id   string
	Data Data
}

type Data struct {
	Value   int     `json:"val"`
	Version float32 `json:"ver"`
	Valid   bool
	Hash    [32]byte
	TxnID   int
}

type details struct {
	Value   int
	Version float32
}
