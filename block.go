package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

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
			var file, _ = os.OpenFile("ledger.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			_, err := file.WriteString(string(content) + "\n")
			file.Close()
			if(err!=nil){
				fmt.Println(err)
			}
		}
		UpdateBlock()
	}
}

func InitChain() {
	fmt.Println("bxhgdcndx")
	block.TimeStamp = time.Now()
	block.CommitStatus = true
	content, err := json.Marshal(block)
	if(err!=nil){
		fmt.Println(err)
	}
	var file, _ = os.OpenFile("ledger.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	_, err = file.WriteString(string(content) + "\n")
	if err != nil {
		fmt.Println("error writing to file")
	}
	fmt.Println("File Created Successfully")
	file.Close()
	UpdateBlock()
}
