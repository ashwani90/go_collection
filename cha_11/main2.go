package main

// import "fmt"

// type Product struct {
// 	name, category string
// 	price float64
// }

// // Defining methods for type alias
// type ProductList []Product

// func (products *ProductList) calcCategoryTotals() map[string]float64 {
// 	totals := make(map[string]float64)
// 	for _, p := range *products {
// 		totals[p.category] = totals[p.category] + p.price

// 	}
// 	return totals
// }

// func printDetails(product *Product) {
// 	fmt.Println("Name: ", product.name)
// }

// // Converting it to a method
// func (product *Product) printDetails() {
// 	fmt.Println("Name: ", product.name, " Price is", product.calcTax(.2,100))
// }

// func (product *Product) calcTax(rate, threshold float64) float64 {
// 	if (product.price > threshold) {
// 		return product.price + (product.price*rate)
// 	}
// 	return product.price
// }

// func main2() {
// 	products := []*Product {
// 		{"Kayak", "Watersports", 275},
// 		{ "LifeJacket", "Watersports", 48.95 },
// 		{"Soccer Ball", "Soccer", 19.50},
// 	}

// 	products2 := ProductList {
// 		{"Kayak", "Watersports", 275},
// 		{ "LifeJacket", "Watersports", 48.95 },
// 		{"Soccer Ball", "Soccer", 19.50},
// 	}

// 	for category, total := range products2.calcCategoryTotals() {
// 		fmt.Println("Category: ", category, " Total:  ", total)
// 	}

// 	for _,p := range products {
// 		p.printDetails()
// 	}
// 	fmt.Println("Hello Methods and interfaces")
// }