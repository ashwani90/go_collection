package main

import (
	"fmt"
	"time"
)


// without go routine
// func CalcStoreTotal(data ProductData) {
// 	var storeTotal float64

// 	for category, group := range data {
// 		storeTotal += group.TotalPrice(category)
// 	}
// 	fmt.Println("Total: ", ToCurrency(storeTotal))
// }

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64)
	for category, group := range data {
		go group.TotalPrice(category, channel)
	}
	fmt.Println("Total: ", ToCurrency(storeTotal))
}

func (group ProductGroup) TotalPrice(category string, resultChannel chan float64) {
	var total float64
	for _,p := range group {
		fmt.Println(category, "product: ", p.Name)
		total += p.Price
		time.Sleep(time.Millisecond*1000)
	}
	fmt.Println(category, "Subtotal:", ToCurrency(total))
	resultChannel<-total
	return
}