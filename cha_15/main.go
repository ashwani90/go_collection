package main

import "fmt"

func main() {
	// fmt.Println("Hello error handling")

	categories := []string {"Watersports", "Chess"}
	for _, cat := range categories {
		total := Products.TotalPrice(cat)
		fmt.Println(cat, "Total:", ToCurrency(total))
	}
}