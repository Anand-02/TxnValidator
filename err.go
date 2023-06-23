package main

import (
	"fmt"
	"log"
)

func Err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ChkErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
