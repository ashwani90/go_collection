package main

import (
	"fmt"
	"packages/store"
)

func main2() {
	product := store.Product {
		Name: "Kayak",
		Category: "Watersports",
		
	}
	fmt.Println("Name: ", product.Name)
	fmt.Println("Category: ", product.Category)
}

func main() {
	product := store.NewProduct("Kayak", "Watersports", 279)

	fmt.Println("Name: ", product.Name)
	fmt.Println("Category: ", product.Category)
	fmt.Println("Price: ", product.Price())
}

