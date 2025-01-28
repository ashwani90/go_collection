package main

import (
	"fmt"
	"cha_13/composition/store"
)

func main() {
	product := store.NewProduct("Kayak", "Watersports", 279)
	deal := store.NewSpecialDeal("Weekend Special", product, 50)
	Name, price, Price := deal.GetDetails()

	fmt.Println("Name: ", Name)
	fmt.Println("Price field: ", price)
	fmt.Println("Price method: ", Price)
}

func main2() {
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

	rentals := []*store.RentalBoat {
		store.NewRentalBoat("Rubber", 20, 2, false, false, "N/A", "N/A"),
		store.NewRentalBoat("yacht", 50000, 2, true, false, "Bob", "Alice"),
		store.NewRentalBoat("Super Yacht", 200, 2, false, false, "Dora", "Charlie"),
		store.NewRentalBoat("Rubber", 20, 2, false, false, "Dora", "Charlie"),
	}

	for _, r := range rentals {
		fmt.Println("Rental Boat:", r.Name, "Rental price: ", r.Price(0.2), "Captain: ", r.Captain)
	}
}
