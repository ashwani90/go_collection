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

func (p Product) getName() string {
	return p.name
}

func (p Product) getCost(_ bool) float64 {
	return p.price
}

func (s Service) getName() string {
	return s.description
}

func (s Service) getCost(recur bool) float64 {
	if (recur) {
		return s.monthlyFee * float64(s.durationMonths)
	}
	return s.monthlyFee
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


	// using the interface
	expenses := []Expense {
		Product {"Kayak", "Watersports", 275},
		Service {"Boat Cover", 12, 89.50},
	}

	for _, expense := range expenses {
		fmt.Println("Expense: ", expense.getName(), " Cost: ", expense.getCost(true))
	}
}