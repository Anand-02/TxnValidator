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
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		panic(fmt.Errorf("could not open leveldb database: %w", err))
	}
	return db
}

func Put(key []byte, value []byte) {
	valueExists, _ := base.Has(key, nil)
	if valueExists {
		Delete(key)
	}
	base.Put(key, value, nil)
}

func Get(key []byte) []byte {
	output, getErr := base.Get(key, nil)
	Err(getErr)
	return output
}

func Delete(key []byte) {
	base.Delete(key, nil)
}
