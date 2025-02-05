package main

import "fmt"

func main() {
	// fmt.Println("Hello error handling")

	categories := []string {"Watersports", "Chess", "Running" }
	for _, cat := range categories {
		total, err := Products.TotalPrice(cat, )
		if (err == nil) {
			fmt.Println(cat, "Total:", ToCurrency(total))
		} else {
			fmt.Println(cat, "(no such category)")
		}
		
	}
}