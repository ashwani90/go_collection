package main

import "fmt"

type Product struct {
	name, category string
	price float64
}

type Supplier struct {
	name, city string
}

func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (product *Product) printDetails() {
	fmt.Println("Name: ", product.name)
}

func (product *Product) calcTax(rate, threshold float64) float64 {
	if (product.price > threshold) {
		return product.price + (product.price*rate)
	}
	
	return product.price
}

func (Supplier *Supplier) printDetails() {
	fmt.Println("Supplier: ", Supplier.name, " City: ", Supplier.city)
}

func main() {
	products := []*Product {
		newProduct("Kayak", "Watersports", 275),
		newProduct("LifeJacket", "Watersports", 48.95),
		newProduct("Soccer Ball", "Soccer", 19.50),
	}

	for _, p := range products {
		p.printDetails()
	}

	suppliers := []*Supplier {
		{"Acme Co", "New York City"},
		{"BoatCo", "Chicago"},
	}

	for _, s := range suppliers {
		s.printDetails()
	}
}
