package main

import (
	"encoding/json"
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
		strData, e := json.Marshal(tmp)
		key := []byte(id)
		ChkErr(e)
		Put(key, strData)
		// temp := Get(key)
		// fmt.Println(string(temp))
	}
	// chkValue(1)
}
