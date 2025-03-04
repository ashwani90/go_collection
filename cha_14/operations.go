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
	// now lets the buffered channe
	// var channel chan float64 = make(chan float64)
	// this is like at a time it will have only two values
	var channel chan float64 = make(chan float64, 2)
	for category, group := range data {
		go group.TotalPrice(category, channel)
	}
	time.Sleep(time.Second*5)
	fmt.Println("-- Starting reading data from the channel")

	for i:=0;i<len(data);i++ {
		fmt.Println("-- channel read pending", len(channel), "items in buffer", cap(channel))
		value := <-channel
		fmt.Println("-- channel read complete", value)
		storeTotal += value
		time.Sleep(time.Second)
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
	fmt.Println(category, "channel sending:", ToCurrency(total))
	resultChannel<-total
	fmt.Println(category, "channel send complete")
	return
}