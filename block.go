package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var file, _ = os.OpenFile("ledger.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

func UpdateBlock() {
	block.DeriveHash()
	block.BlockNo++
	block.PrevBlockHash = LastHash
	block.Txns = make([]Txn, 0)
	for len(block.Txns) < BlockCapacity {
		block.Txns = append(block.Txns, <-ch)
	}
	if len(block.Txns) == BlockCapacity {
		block.TimeStamp = time.Now()
		block.CommitStatus = true
		content, err := json.Marshal(block)
		if err != nil {
			fmt.Println("error while writing block")
		} else {
			_, err := file.WriteString(string(content) + "\n")
			ChkErr(err)
		}
		UpdateBlock()
	}
}

func InitChain() {
	fmt.Println("bxhgdcndx")
	block.TimeStamp = time.Now()
	block.CommitStatus = true
	content, err := json.Marshal(block)
	ChkErr(err)
	_, err = file.WriteString(string(content) + "\n")
	if err != nil {
		fmt.Println("error writing to file")
	}
	fmt.Println("File Created Successfully")
	UpdateBlock()
}
