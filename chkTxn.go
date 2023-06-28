package main

import (
	"encoding/json"
	"fmt"
)

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
		go newTxn.DeriveHash(tmp)
		strData, err := json.Marshal(tmp)
		if err != nil {
			fmt.Println(err)
		}
		Put(key, strData)
	}
}
