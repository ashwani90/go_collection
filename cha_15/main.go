package main

import "fmt"

func main() {
	recoveryFunc := func() {
		if arg := recover(); arg != nil {
			if err, ok := arg.(error); ok {
				fmt.Println("Error:", err.Error())
			} else if str, ok := arg.(string);ok {
				fmt.Println("Message:", str)
			} else {
				fmt.Println("Panic Recovered")
			}
		}
	}
	defer recoveryFunc()
	categories := []string {"Watersports", "Chess", "Running" }
	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if (message.CategoryError == nil) {
			fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
		} else {
			// Dealing with unrecoverable errors
			// fmt.Println(message.Category, "(no such category)")
			panic(message.CategoryError)
		}
	}
}

func main2() {
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