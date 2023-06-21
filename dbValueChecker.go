package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func chkValue(i int) {
	id := "SIM" + strconv.Itoa(i)
	key, _ := json.Marshal(id)
	data := Get(key)
	var tmp details
	json.Unmarshal(data, &tmp)
	fmt.Println(tmp)
}
