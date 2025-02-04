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

// func enumerateProducts(channel chan<- *Product) {
// 	for _,p := range ProductList[:3] {
// 		channel <- p
// 		time.Sleep(time.Millisecond*800)
// 	}
// 	close(channel)
// }

func enumerateProducts2(channel chan<- *Product) {
	for _, p := range ProductList {
		select {
		case channel<-p:
			fmt.Println("Sent product:", p.Name)
		default:
			fmt.Println("Discarding product:", p.Name)
			time.Sleep(time.Second)
		}
	}
	close(channel)
}
// sending to multiple channels
func enumerateProducts2(channel1, channel2 chan<- *Product) {
	for _, p := range ProductList {
		select {
		case channel1<-p:
			fmt.Println("Sent via channel 1:", p.Name)
		case channel2<-p:
			fmt.Println("Sent via channel 2:", p.Name)
		default:
			fmt.Println("Discarding product:", p.Name)
			time.Sleep(time.Second)
		}
	}
	close(channel1)
	close(channel2)
}

func main() {
	// productChannel := make(chan *Product, 5)
	// go enumerateProducts(productChannel)
	// time.Sleep(time.Second)
	// for p := range productChannel {
	// 	fmt.Println("Received product:", p.Name)
	// }

	c1 := make(chan *Product, 2)
	c2 := make(chan *Product, 2)
	go enumerateProducts(c1, c2)
	time.Sleep(time.Second)
	for p := range c1 {
		fmt.Println("Channel 1 received product:", p.Name)
	}
	for p := range c2 {
		fmt.Println("Channel 2 received product:", p.Name)
	}
}

func main3() {
	dispatchChannel := make(chan DispatchNotification, 100)
	// var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
	// var receiveChannel <-chan DispatchNotification = dispatchChannel

	productChannel := make(chan *Product)
	go enumerateProducts(productChannel)
	openChannels := 2

	// type conversion we can alos do
	// go DispatchOrders(sendOnlyChannel)
	// receiveDispatches(receiveChannel)

	// go DispatchOrders(chan<- DispatchNotification(dispatchChannel))
	// receiveDispatches((<-chan DispatchNotification)(dispatchChannel))

	for {
		select {
		case details, ok:= <- dispatchChannel:
			if ok {
				fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, 'x', details.Product.Name)
			} else {
				fmt.Println("Channel has been closed")
				dispatchChannel = nil
				openChannels--
				// goto alldone
			}
		case product, ok := <- productChannel:
			if ok {
				fmt.Println("Product", product.Name)
			} else {
				fmt.Println("Product channel has been closed")
				productChannel = nil
				openChannels--
			}
		default:
			if (openChannels == 0) {
				goto alldone
			}
			fmt.Println("-- No message ready to be received")
			time.Sleep(time.Millisecond*500)
		}
	}
	alldone: fmt.Println("All values received")
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