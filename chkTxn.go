package main

import (
	"encoding/json"
	//"fmt"
	"sync"
)

var NewTxnMutex = &sync.Mutex{}


func Validator(inputTxn map[string]Data) {
	TxnNo++
	newTxn := Txn{}
	for id, value := range inputTxn {
		newTxn.Id = id
		newTxn.Data = value
		key := []byte(id)
		data := Get(key)
		var tmp details
		json.Unmarshal(data, &tmp)
		if tmp.Version == newTxn.Data.Version {
			newTxn.Data.Valid = true
			newTxn.Data.Value = value.Value
			newTxn.Data.Version += 1.0
			tmp.Version += 1.0
			tmp.Value = value.Value
		}
		newTxn.Data.TxnID = TxnNo
		//NewTxnMutex.Lock()
		go newTxn.DeriveHash(tmp)
		//NewTxnMutex.Unlock()
		//fmt.Println(newTxn)
		strData, e := json.Marshal(tmp)
		ChkErr(e)
		Put(key, strData)
		// temp := Get(key)
		// fmt.Println(string(temp))
	}
	
}
