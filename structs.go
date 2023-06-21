package main

import "time"

type Block struct {
	BlockNo       int
	PrevBlockHash []byte
	Txns          []Txn
	TimeStamp     time.Time
}

type Txn struct {
	Id   string
	Data Data
}

type Data struct {
	Value   int     `json:"val"`
	Version float32 `json:"ver"`
	Valid   bool
}

type details struct {
	Value   int
	Version float32
}
