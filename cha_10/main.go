package main

import (
	"fmt"
	"encoding/json"
	"strings"
)

func writeName(val struct { 
	name, category string
	price float64
 }) {
	fmt.Println("Name: ", val.name)
 }

func parsing() {
	type Prod2 struct {
		name, category string
		price float64
	}

	prod := Prod2 { name: "Kayak", category: "Watersports", price: 275.00 }
	var builder strings.Builder
	json.NewEncoder(&builder).Encode(struct {
		ProductName string
		ProductPrice float64
	}{
		ProductName: prod.name,
		ProductPrice: prod.price,
	})
	fmt.Println(builder.String())
}

func main() {

	type Product struct {
		name, category string
		price float64
	}

	type StockLevel struct {
		Product
		Alternate Product
		count int
	}

	p1 := Product {
		name: "Kayak",
		category: "Watersports",
		price: 275,
	}

	p2 := p1
	p1.name = "Original Kayak"
	fmt.Println("P1: ", p1.name)
	fmt.Println("P2: ", p2.name)

	p3 := &p1
	fmt.Println("P3: ", (*p3).name)

	array := [1] StockLevel {
		{
			Product: Product { "Kayak", "Watersports", 275.00 },
			Alternate: Product { "Lifejacket", "Watersports", 48.95 },
			count: 100,
		},
	}
	fmt.Println("Array: ", array[0].Product.name)
	slice := [] StockLevel {
		{
			Product: Product { "Kayak", "Watersports", 275.00 },
			Alternate: Product { "Lifejacket", "Watersports", 48.87 },
			count: 100,
		},
	}
	fmt.Println("Slice: ", slice[0].Product.name)

	kvp := map[string]StockLevel {
		"kayak": {
			Product: Product { "Kayak", "Watersports", 275.00},
			Alternate: Product {"Lifejacket", "Watersorts", 28.95},
			count: 100,
		},
	}

	fmt.Println("Map: ", kvp["kayak"].Product.name)

	kayak := Product {
		name: "Kayak",
		category: "Watersports",
		// price: 275,
	}

	stockItem := StockLevel {
		Product: Product { "Kayak", "Watersports", 275.00},
		Alternate: Product {"Lifejacket", "Watersports", 49.89 },
		count: 100,
	}

	fmt.Println("Name: ", stockItem.Product.name)
	fmt.Println("Count: ", stockItem.count)
	fmt.Println("Alt name: ", stockItem.Alternate.name)

	fmt.Println(kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	fmt.Println("Changed price: ", kayak.price)

	var lifejacket Product

	fmt.Println("Name is 0 value", lifejacket.name == "")
	fmt.Println("category is 0 value", lifejacket.category == "")
	fmt.Println("price is 0 value", lifejacket.price == 0)

	// using literal syntax
	var kaya = Product { "Kayak", "Watersports", 275.0 }
	fmt.Println("Name: ", kaya.name)
	fmt.Println("Category: ", kaya.category)
	fmt.Println("Price: ", kaya.price)

	fmt.Println("hello structs")


	// Converting between sturct types
	type Item struct {
		name string
		category string
		price float64
	}

	prod := Product {name: "Kayak", category: "Watersports", price: 275.00}
	item := Product { name: "Kayak", category: "Watersports", price: 275.00}
	fmt.Println("prod == item", prod == Product(item))
	// Kinf of polymorphism
	writeName(prod)
	writeName(item)

	parsing()
}