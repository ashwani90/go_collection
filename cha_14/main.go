package main

import (
	"fmt"
	"time"
)

func receiveDispatches(channel <-chan DispatchNotification) {
	for details := range channel {
		fmt.Println("Dispatch to", details.Customer, ";", details.Quantity, "x", details.Product.Name)
	}
	fmt.Println("Channel has been closed")
}

// use go run . to run the project


func main() {
	dispatchChannel := make(chan DispatchNotification, 100)
	var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
	var receiveChannel <-chan DispatchNotification = dispatchChannel

	// type conversion we can alos do
	go DispatchOrders(sendOnlyChannel)
	receiveDispatches(receiveChannel)

	go DispatchOrders(chan<- DispatchNotification(dispatchChannel))
	receiveDispatches((<-chan DispatchNotification)(dispatchChannel))
}

func main2() {
	fmt.Println("main function started")
	CalcStoreTotal(Products)
	time.Sleep(time.Second*5)
	fmt.Println("main function ended")

	dispatchChannel := make(chan DispatchNotification, 100)
	go DispatchOrders(dispatchChannel)
	for {
		if details, open := <- dispatchChannel; open {
			fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
		} else {
			fmt.Println("Channel has been closed")
			break
		}
	}

	for details := range dispatchChannel {
		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
	}
	fmt.Println("Channel has been closed")
	// this gives error that deadlock
}