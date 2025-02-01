package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DispatchNotification struct {
	Customer string
	*Product
	Quantity int
}

var Customers = []string{"Alice", "Bob", "Happu", "Vibhuti", "Manmohan", "Tillu"}

// channle value will be of type DispatchNotification
// decides on the direction so this in other way can not be done
func DispatchOrders(channel chan<- DispatchNotification) {
	rand.Seed(time.Now().UTC().UnixNano())
	orderCount := rand.Intn(3)+2
	fmt.Println("OrderCount: ", orderCount)
	for i:=0;i<orderCount;i++ {
		channel <- DispatchNotification{
			Customer: Customers[rand.Intn(len(Customers)-1)],
			Quantity: rand.Intn(10),
			Product: ProductList[rand.Intn(len(ProductList)-1)],
		}
		if (i==1) {
			notification := <- channel
			fmt.Println("Read:", notification.Customer)
		}
	}
	close(channel)
}