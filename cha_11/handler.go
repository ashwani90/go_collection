package main

import (
	"fmt"
	./Product
	./Service
)

func main() {
	kayak := Product {"Kayak", "Watersports", 275}
	insurance := Service {"Boat Cover", 12, 989.50}

	fmt.Println("product: ", kayak.name, "Price: ", kayak.price)
	fmt.Println("Service: ", insurance.description, "Price: ", insurance.monthlyFee*float64(insurance.durationMonths))
}