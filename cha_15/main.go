package main

import "fmt"

func (slice ProductSlice) TotalPrice(category string) (total float64) {
	for _, p := range slice {
		if (p.Category == category) {
			total += p.Price
		}
	}
	return
}

func main() {
	// fmt.Println("Hello error handling")

	categories := []string {"Watersports", "Chess"}
	for _, cat := range categories {
		total := Products.TotalPrice(cat)
		fmt.Println(cat, "Total:", ToCurrency(total))
	}
}