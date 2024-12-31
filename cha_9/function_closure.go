package main

import "fmt"

type calcFunc func(float64) float64

func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product: ", product, " Price is: ", calculator(price))
} 

func main() {
	waterSportsProducts := map[string] float64 {
		"keyak": 275,
		"Life Jacket": 48.95,
	}

	soccerProducts := map[string] float64 {
		"Soccer Ball": 19.50,
		"Stadium": 79500,
	}

	calc := func(price float64) float64 {
		if (price > 100) {
			return price + (price*0.2)
		}
		return price
	}

	for product, price := range waterSportsProducts {
		printPrice(product, price, calc)
	}

	calc = func(price float64) float64 {
		if (price>50) {
			return price + (price*0.1)
		}
		return price
	}

	for product, price := range soccerProducts {
		printPrice(product, price, calc)
	}
}