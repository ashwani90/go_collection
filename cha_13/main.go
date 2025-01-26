package main

import (
	"fmt"
	"cha_13/composition/store"
)

func main() {
	fmt.Println("Hello type and interface composition")
	kayak := store.NewProduct("Kayak", "Watersports", 275)
	lifejacket := &store.Product { Name: "Lifejacket", Category: "Watersports"}

	for _, p := range []*store.Product { kayak, lifejacket} {
		fmt.Println("Name: ", p.Name, "Category: ", p.Category, "Price: ", p.Price(0.2))
	}
}
