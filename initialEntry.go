package main

import (
	"encoding/json"
	"fmt"
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
	chkValue(100)
}
