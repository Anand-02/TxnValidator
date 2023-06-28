package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var now time.Time

func myfunc() {
	for {
		select {
		case v := <-tmpch:
			ch <- v

			if len(ch) == BlockCapacity {
				now = time.Now()
				UpdateBlock()
			}
		case <-time.After(5 * time.Second):
			if len(ch) > 0 {
				now = time.Now()
				UpdateBlock()
			}
		}
	}
}

func UpdateBlock() {
	block.DeriveHash()
	block.BlockNo++
	block.PrevBlockHash = LastHash
	block.Txns = make([]Txn, 0)
	for len(ch) > 0 {
		block.Txns = append(block.Txns, <-ch)
	}
	if len(block.Txns) > 0 {
		block.TimeStamp = time.Now()
		block.CommitStatus = true
		content, err := json.Marshal(block)
		if err != nil {
			fmt.Println("error while writing block")
		} else {
			var file, _ = os.OpenFile("ledger.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			_, err := file.WriteString(string(content) + "\n")
			fmt.Println("Processing time of Block No", block.BlockNo, time.Since(now))
			file.Close()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func InitChain() {
	block.TimeStamp = time.Now()
	block.CommitStatus = true
	content, err := json.Marshal(block)
	if err != nil {
		fmt.Println(err)
	}
	var file, _ = os.OpenFile("ledger.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	_, err = file.WriteString(string(content) + "\n")
	if err != nil {
		fmt.Println("error writing to file")
	}
	fmt.Println("File Created Successfully")
	file.Close()
	go myfunc()
}
