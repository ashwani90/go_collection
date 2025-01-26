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

	boats := []*store.Boat {
		store.NewBoat("Kayak", 275, 1, false),
		store.NewBoat("Canoe", 400, 2, false),
		store.NewBoat("Tender", 775, 3, true),
	}

	for _, b := range boats {
		fmt.Println("Conventional: ", b.Product.Name, "Direct: ", b.Name)
	}
}
