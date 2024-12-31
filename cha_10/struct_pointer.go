package main

import "fmt"

type Product struct {
	name, category string
	price float64
	*Supplier
}

type Supplier struct {
	name, city string
}

// func newProduct(name category string, price float64, supplier *Supplier) *Product {
// 	return &Product{name, category, price-10, supplier}
// }

func calcTax(product *Product) {
	if ((*product).price > 100) {
		(*product).price += (*product).price * .2

	}

}

func calcTax2(product *Product) *Product {
	if (product.price > 100) {
		product.price += product.price * .2
	}
	return product
}

// creating struct constructor function
func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	return &Product{name, category, price, supplier}
}

func copyProduct(product *Product) Product {
	p := *product
	s := *product.Supplier
	p.Supplier = &s 
	return p
}

func main() {
	
	kayak := Product {
		name: "Kayak",
		category: "Watersports",
		price: 275,
	}
	kayak2 := &Product {
		name: "Kayak",
		category: "Watersports",
		price: 275,
	}
	
	calcTax(&kayak)
	fmt.Println("Price: ", kayak.price)
	calcTax2(&kayak)
	fmt.Println("Name: ", kayak.name)
	fmt.Println("Price: ", kayak.price)
	calcTax(kayak2)
	fmt.Println("Price: ", kayak2.price)
	calcTax2(kayak2)
	fmt.Println("Name: ", kayak2.name)
	fmt.Println("Price: ", kayak2.price)
	kayak2 = calcTax2(&Product {
		name: "Kayak",
		category: "Watersports",
		price: 275,
	})
	fmt.Println("Name: ", kayak2.name, " Category: ", kayak2.category, " Price: ", kayak2.price)

	// products := [2]*Product {
	// 	newProduct("Kek", "aad", 23.2),
	// 	newProduct("Ke2k", "aad2", 23.22),
	// }

	// for _, p := range products {
	// 	fmt.Println(p.name)
	// }

	// using pointer type for struct fields
	acme := &Supplier {"ace", "New"}
	products := [2]*Product {
		newProduct("Kek", "aad", 23.2, acme),
		newProduct("Ke2k", "aad2", 23.22, acme),
	}
	p1 := newProduct("Kek", "aad", 23.2, acme)
	p2 := copyProduct(p1)
	fmt.Println(p2.name)

	for _, p := range products {
		fmt.Println(p.name)
	}
	var prod Product
	var prodPtr *Product
	fmt.Println("Pointer", prodPtr)
	fmt.Println("Pointer", prod)
}