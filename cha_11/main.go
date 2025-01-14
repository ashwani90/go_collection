package main

import (
	"fmt"
)

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

// type Product struct {
// 	name, category string
// 	price float64
// }

// func (product *Product) getName() string {
// 	fmt.Println("Name: ", product.name)
// 	return "a"
// }

// func (product *Product) getCost(annual bool) float64 {
// 	fmt.Println("Name: ", product.name)
// 	return 25
// }

func main() {
	product := Product {"Kayak", "Watersports", 275}

	var expense Expense = &product
	product.price = 100

	fmt.Println("Product fiueld type: ", product.price)
	fmt.Println("Product fiueld type: ", expense.getName())

	expenses := []Expense {
		Service {"Boat Cover", 12,89.50, []string{}},
		Service {"Boat Cover 2", 12,89.50, []string{}},
	}

	// this will fail in type conversion
	for _, expense := range expenses {
		s := expense.(Service)
		fmt.Println("Service: ", s.description)
	}

	expenses = []Expense {
		Service {"Boat Cover", 12,89.50, []string{}},
		Service {"Boat Cover 2", 12,89.50, []string{}},
		&Product {"Boat Cover 2", "Water",89.50},
	}

	for _, expense := range expenses {
		if s, ok := expense.(Service);ok {
			fmt.Println("Service: ", s.description)
		} else {
			fmt.Println("Unable to convert ", s.getName())
		}
	}

	// using dynamic types
	
}