package main

import (
	"fmt"
)

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

type Person struct {
	name, city string
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
	for _, expense := range expenses {
		switch value := expense.(type) {
		case Service:
			fmt.Println("Type is Service", value.description)
		case *Product:
			fmt.Println("Type is pointer to Product")
		default:
			fmt.Println("Something other")
		}
	}

	// using empty interface
	data := []interface{} {
		expense,
		Product {"LifeJacket", "Wa", 12},
		Service {"Bot", 12, 89, []string{}},
		Person {"Alice", "London"},
		&Person {"Bob", "New York"},
		"This is a string",
		100,
		true,
	}
	for _, item := range data {
        switch value := item.(type) {
            case Product:
                fmt.Println("Product:", value.name, "Price:", value.price)           
            case *Product:
                fmt.Println("Product Pointer:", value.name, "Price:", value.price)                           
            case Service:
                fmt.Println("Service:", value.description, "Price:", 
                    value.monthlyFee * float64(value.durationMonths))
            case Person:                
                fmt.Println("Person:", value.name, "City:", value.city)
            case *Person:
                fmt.Println("Person Pointer:", value.name, "City:", value.city)                
            case string, bool, int:
                fmt.Println("Built-in type:", value)
            default:
                fmt.Println("Default:", value)       
        }
    }
	
}