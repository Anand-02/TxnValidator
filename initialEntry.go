package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func SetKeys() {
	defer base.Close()
	for i := 1; i <= 1000; i++ {
		id := fmt.Sprintf("SIM%d", i)
		tmp := details{
			Value:   i,
			Version: 1.0,
		}
		strData, e := json.Marshal(tmp)
		key := []byte(id)
		ChkErr(e)
		Put(key, strData)
	}
	id := "SIM" + strconv.Itoa(100)
	key := []byte(id)
	data := Get(key)
	var tmp details
	json.Unmarshal(data, &tmp)
	fmt.Println(tmp)
}
