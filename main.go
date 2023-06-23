package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// var db = openDB()
func (t *details) DeriveHash() [32]uint8 {
	info, _ := json.Marshal(t)
	hash := sha256.Sum256([]byte(info))
	fmt.Printf("%T", hash)
	return hash
}

func main() {
	// defer openDB().Close()
	SetKeys()
	createRoute()
}
