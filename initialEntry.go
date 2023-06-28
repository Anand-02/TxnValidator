package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func SetKeys() {
	//defer base.Close()
	for i := 1; i <= 1000; i++ {
		id := "SIM" + strconv.Itoa(i)
		//fmt.Println(id)
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
		// temp := Get(key)
		// fmt.Println(string(temp))
	}
	// chkValue(1)
}
