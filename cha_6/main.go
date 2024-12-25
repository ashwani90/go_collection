package main

import (
	"fmt"
	"strconv"
)

func main() {
	kayakPrice := 275.00

	// usage of scoped variables
	if (kayakPrice > 100 && kayakPrice <= 500) {
		secondVar := 500
		print(secondVar)
		fmt.Println("price is greater than 100 and less than 500")
	} else if (kayakPrice > 500) {
		secondVar := "Price ok"
		print(secondVar)
		fmt.Println("Price greater than 500")
	} else {
		secondVar := false
		print(secondVar)
		fmt.Println("I dont want to classify the price now")
	}

	// using initialization statement with if condition
	// priceString := "275"
	priceString := "a"
	if priceInt,err := strconv.Atoi(priceString); err == nil {
		fmt.Println("Price is : ", priceInt)
	} else {
		fmt.Println("Error is: ", err)
	}

	// using loops
	count := 0
	for {
		fmt.Println("counter: ", count)
		count++
		if (count>3) {
			break
		}
	}

	// now condition inside the loop
	counter := 0
	for (counter<=3) {
		fmt.Println("count value is: ", counter)
		counter++
		// break condition is not needed now
	}

	// now complete for loop
	// count2 := 0
	for count2:=0; count2<=3; count2++ {
		// Continuing loops
		if count2 == 1 {
			continue
		}
		fmt.Println("Count value is: ", count2)
	}

	product := "Kayak"
	// and using switch statement in loop
	for index, character := range product {
		switch (character) {
		case 'K', 'k':
			fmt.Println("K at position: ", index)
			// fallthrough to the next statement
			
			// terminate case statement
			if (character == 'k') {
				fmt.Print("Need to break now")
				break
			}
			fallthrough
		case 'y':
			fmt.Println("y at piosition: ", index)
		default:
			fmt.Println("Inside the default case")
		}
		fmt.Println("Index is: ", index, " Character is: ", character)
	}
	// or you can also loop through the index only
	for index := range product {
		fmt.Println("Index is: ", index)
	}
	// or only through the character only
	for _, character := range product {
		fmt.Println("character is: ", string(character))
	}

	// Enumerating data structures
	
	products := []string { "Abcd", "Efgh", "Ijkl", "Mnop", "Qrst" }
	for _,n := range products {
		fmt.Println("slice is ", n)
	}

	// using initialization method
	for countz := 0;countz<20;countz++ {
		switch(countz/2) {
		case 2,3,5,7:
			fmt.Println("Prime value is: ", countz/2)
		default:
			fmt.Println("Non prime value:", countz/2)
		}
	}

	// optimized version
	for county := 0;county<20;county++ {
		switch va:=county/2;va {
		case 2,3,5,7:
			fmt.Println("Prime value is: ", va)
		default:
			fmt.Println("Non prime value:", va)
		}
	}

	// using label statements
	countx := 0
	target: fmt.Println("counter ", countx)
	countx++
	if countx<5 {
		goto target
	}
}

