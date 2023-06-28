package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func SetKeys() {
	for i := 1; i <= 1000; i++ {
		id := "SIM" + strconv.Itoa(i)
		tmp := details{
			Value:   i,
			Version: 1.0,
		}
		strData, err := json.Marshal(tmp)
		key := []byte(id)
		if err != nil {
			fmt.Println(err)
		}
		Put(key, strData)
	}
}
