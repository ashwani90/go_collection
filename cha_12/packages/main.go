// package main

// import (
// 	"fmt"
// 	"packages/store"
// 	// using package alias
// 	// currencyFmt "packages/fmt"
// 	// we can also use . import
// 	. "packages/fmt"
// 	"packages/store/cart"
// )

// func main2() {
// 	product := store.Product {
// 		Name: "Kayak",
// 		Category: "Watersports",
		
// 	}
// 	fmt.Println("Name: ", product.Name)
// 	fmt.Println("Category: ", product.Category)
// }

// func main() {
// 	product := store.NewProduct("Kayak", "Watersports", 279)

// 	fmt.Println("Name: ", product.Name)
// 	fmt.Println("Category: ", product.Category)
// 	fmt.Println("Price: ", product.Price())
// 	// fmt.Println("Price: ", currencyFmt.ToCurrency(product.Price()))
// 	fmt.Println("Price: ", ToCurrency(product.Price()))
// 	cart := cart.Cart {
// 		CustomerName: "Alice",
// 		Products: []store.Product{*product},
// 	}
// 	fmt.Println("Name: ", cart.CustomerName)
// 	fmt.Println("Total: ", ToCurrency(cart.GetTotal()))
// }

package main 

import (
    "fmt"
    "packages/store"
    . "packages/fmt"
    "packages/store/cart"
    _ "packages/data"
    //"github.com/fatih/color"
)

func main() {

    product := store.NewProduct("Kayak", "Watersports", 279) 

    cart := cart.Cart {
        CustomerName: "Alice",
        Products: []store.Product{ *product },
    }

    // color.Green("Name: " + cart.CustomerName)
    // color.Cyan("Total: " + ToCurrency(cart.GetTotal()))
    fmt.Println("Name:", cart.CustomerName)
    fmt.Println("Total:",  ToCurrency(cart.GetTotal()))
}