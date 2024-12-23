package main

import "fmt"

func main() {
	fmt.Println("Hello Go!")
	for i:=0;i<5;i++ {
		fmt.Println(i)
	}
}

// To debug this code
// dlv debug main.go
// create new break point break bp1 main.main:3
// continue - continue executing
// this can also be integrated into editor