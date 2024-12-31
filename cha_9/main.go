package main

import "fmt"

func calcWithTax(price float64) float64 {
	return price + (price*.02)
}

func calcWithoutTax(price float64) float64 {
	return price
}

// Passing function as argument
func printPrice(product string, price float64, calculator func(float64) float64) {
	fmt.Println("Product: ", product, " Price is: ", calculator(price))
}

// using function as result value
func selectCalculator(price float64) func(float64) float64 {
	if (price > 100) {
		return calcWithTax
	}
	return calcWithoutTax
}

// creating function type aliases
type calcFunc func(float64) float64

// Passing function as argument
func printPrice2(product string, price float64, calculator calcFunc) {
	fmt.Println("Product: ", product, " Price is: ", calculator(price))
}

// using literal function syntax
func selectCalculator2(price float64) calcFunc {
	if (price > 100) {
		var withTax calcFunc = func (price float64) float64 {
			return price + (price*0.2)
		}
		return withTax
	}
	withoutTax := func (price float64) float64 {
		return price
	}
	return withoutTax
}

// using function values directly
func selectCalculator3(price float64) calcFunc {
	if (price > 100) {
		return func (price float64) float64 {
			return price + (price*0.2)
		}
	}
	return func (price float64) float64 {
		return price
	}
}

func main() {
	products := map[string] float64 {
		"Kayak": 275, 
		"LifeJacket": 48.95,
	}
	for product, price := range products {
		printPrice(product, price, selectCalculator(price))
		printPrice2(product, price, selectCalculator3(price))
	}

	
}

func main2() {

	products := map[string] float64 {
		"Kayak": 275, 
		"LifeJacket": 48.95,
	}

	for product, price := range products {
		var calcFunc func(float64) float64

		// func value comparison
		fmt.Println("Function assigned: ", calcFunc==nil)
		if (price > 200) {
			calcFunc = calcWithTax
		} else {
			calcFunc = calcWithoutTax
		}
		printPrice(product, price, calcFunc)

		// fmt.Println("Function assigned: ", calcFunc==nil)
		// totalPrice := calcFunc(price)
		// fmt.Println("Product: ", product, " Price: ", totalPrice)
	}
	fmt.Println("Hello")
}