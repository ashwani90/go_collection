package main


import "fmt"

type CategoryCountMessage struct {
	Category string
	Count int
	TerminalError interface{}
}

func processCategories(categories []string, outChan chan <- CategoryCountMessage) {
	defer func() {
		if arg := recover(); arg != nil {
			fmt.Println(arg)
			outChan <- CategoryCountMessage {
				TerminalError: arg,
			}
			close(outChan)
		}
	}()
	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			outChan <- CategoryCountMessage {
				Category: message.Category,
				Count: int(message.Total),
			}
		} else {
			panic(message.CategoryError)
		}
	}
	close(outChan)
}

func main() {
	categories := []string {"Watersports", "Chess", "Running" }
	channel := make(chan CategoryCountMessage)
	go processCategories(categories, channel)
	for message := range channel {
		if (message.TerminalError == nil) {
			fmt.Println(message.Category, "Total:", message.Count)
		} else {
			fmt.Println("A Terminal error has occures")
		}
		
	}
}

func main3() {

	
	// recoveryFunc := func() {
	// 	if arg := recover(); arg != nil {
	// 		if err, ok := arg.(error); ok {
	// 			fmt.Println("Error:", err.Error())
	// 		} else if str, ok := arg.(string);ok {
	// 			fmt.Println("Message:", str)
	// 		} else {
	// 			fmt.Println("Panic Recovered")
	// 		}
	// 	}
	// }

	// can also use anonymous function
	defer func() {
		if arg := recover(); arg != nil {
			if err, ok := arg.(error); ok {
				fmt.Println("Error:", err.Error())
				// panic after recovering
				panic(err)
			} else if str, ok := arg.(string);ok {
				fmt.Println("Message:", str)
			} else {
				fmt.Println("Panic Recovered")
			}
		}
	}()
	// defer recoveryFunc()
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