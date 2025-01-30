package main

import (
	"fmt"
	"time"
)

// use go run . to run the project

func main() {
	fmt.Println("main function started")
	CalcStoreTotal(Products)
	time.Sleep(time.Second*5)
	fmt.Println("main function ended")
}