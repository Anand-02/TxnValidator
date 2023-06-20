package main

import "fmt"

func Err(err error) {
	if err != nil {
		panic(err)
	}
}

func ChkErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
