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

	// this gives error of ambiguous selector bundle.Price
	// type OfferBundle struct {
	// 	*store.SpecialDeal
	// 	*store.Product
	// }

	// bundle := OfferBundle {
	// 	store.NewSpecialDeal("Weekend Special", product, 50),
	// 	store.NewProduct("LifeJacket", "Watersorts", 48.95),
	// }

	// fmt.Println("Price: ", bundle.Price(0))


	// this in current state will give error
	// to solve this we will use the composition feature
	// products := map[string]*store.Product {
	// 	"kayak": store.NewBoat("Kayak", 279, 1, false),
	// 	"Ball": store.NewProduct("Soccer Ball", "Soccer", 19.50),
	// }

	// for _, p := range products {
	// 	fmt.Println("Name: ", p.Name, "Category: ", p.Category, "Price: ", p.Price(0.2))
	// }

	products := map[string]*store.Product {
		"kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball": store.NewProduct("Soccer Ball", "Soccer", 19.50),
	}

	for _, p := range products {
		fmt.Println("Name: ", p.Name, "Category: ", p.Category, "Price: ", p.Price(0.2))
	}
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
