package main

import (
	"fmt"
	"path"
	"runtime"

	"github.com/syndtr/goleveldb/leveldb"
)

var base = openDB()

func openDB() *leveldb.DB {
	_, filename, _, _ := runtime.Caller(0)
	dbPath := path.Dir(filename) + "/leveldb"
	d, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		panic(fmt.Errorf("could not open leveldb database: %w", err))
	}
	return d
}

func Put(key []byte, value []byte) {
	base.Put(key, value, nil)
}

func Get(key []byte) []byte {
	output, getErr := base.Get(key, nil)
	if getErr != nil {
		fmt.Println(getErr.Error() + string(key))
	}

	return output
}

func Delete(key []byte) {
	base.Delete(key, nil)
}
