package main

// import "fmt"

// type Product struct {
// 	name, category string
// 	price float64
// }

// func printDetails(product *Product) {
// 	fmt.Println("Name: ", product.name)
// }

// // Converting it to a method
// func (product Product) printDetails() {
// 	fmt.Println("Name: ", product.name, " Price is", product.calcTax(.2,100))
// }

// func (product Product) calcTax(rate, threshold float64) float64 {
// 	if (product.price > threshold) {
// 		return product.price + (product.price*rate)
// 	}
// 	return product.price
// }

// func using_pointer() {
// 	kayak := &Product { "Kayak", "Watersports", 275 }
// 	kayak.printDetails()
// }