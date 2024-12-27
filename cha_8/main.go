package main

import "fmt"

// declaring a simple function
func printPrice(price float64, kayak string) {
	fmt.Println(price, kayak)
}

func printSuppliers(suppName string, suppliers ...string) {
	for _, supplier := range suppliers {
		fmt.Println(supplier)
	}
}

func swapValues(first, second *int) {
	fmt.Println("Before swap: ", *first, *second)
	temp := *first
	*first = *second
	*second = temp
	fmt.Println("Before swap: ", *first, *second)
}

// We can modify it to return two variables
func calcTax(price float64) (float64, bool) {
	if (price>100) {
		return price * 0.2, true
	}
	return 0,false
}

// using named results
func calcTotalPrice(products map[string]float64, minSpend float64) (total, tax float64) {
	// using defer call
	fmt.Println("Function started")
	defer fmt.Println("First defer call")
	total = minSpend
	for _, price := range products {
		if taxAmount, due := calcTax(price);due {
			total += taxAmount
			tax += taxAmount
		} else {
			total += price
		}
	}
	defer fmt.Println("Second defer call")
	return
}

func main() {
	fmt.Println("Hello functions")
	kayakPrice := 275.4
	kayakString := "Kayak"
	printPrice(kayakPrice, kayakString)

	// using variadic parameters
	printSuppliers(kayakString, "one", "two", "three")
	// omitting variadic parameters
	printSuppliers(kayakString)

	// using slice as values of variadic parameters
	names := []string {"Hello", "Uno", "Dos", "Tres"}
	printSuppliers(kayakString, names...)

	// swap values using pointers
	val1, val2 := 10, 20
	fmt.Println("before calling function ", val1, val2)
	swapValues(&val1, &val2)
	fmt.Println("after calling function ", val1, val2)

	// Defining and using function results
	products := map[string]float64 {
		"Kayak": 275, 
		"Lifejacket": 49,
	}

	for product, price := range products {
		priceWithTax, taxDue := calcTax(price)
		if (taxDue) {
			fmt.Println("product: ", product, "price: ", priceWithTax)
		} else {
			fmt.Println("product: ", product, "price: ", priceWithTax, "No tax due")
		}
		
	}

	// using named results
	total1, tax1 := calcTotalPrice(products, 10)
	fmt.Println("Total 1: ", total1, "Tax 1: ", tax1)
}