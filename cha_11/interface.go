package main

import (
	"fmt"
)

type Service struct {
	description string
	durationMonths int
	monthlyFee float64
}

type Product struct {
	name, category string
	price float64
}

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func main() {
	kayak := Product {"Kayak", "Watersports", 275}
	insurance := Service {"Boat Cover", 12, 89.40}

	fmt.Println("Product: ", kayak.name, "Price: ", kayak.price)
	fmt.Println("Service: ", insurance.description, "Price: ", insurance.monthlyFee*float64(insurance.durationMonths))
}