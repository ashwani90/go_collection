package main


import "fmt"

type Product struct {
	Name, Category string
	Price float64
}

var Kayak = Product{
	Name: "Kayak",
	Category: "Watersports",
	Price: 275,
}

var Products =[]Product {
	{"Keyak", "watersports", 279},
	{"Lifejacket", "watersports", 279},
	{"Soccer Ball", "Soccer", 279},
	{"Corner Flags", "Soccer", 279},
	{"Stadium", "Soccer", 279},
	{"Thinking Cap", "Chess", 279},
	{"Unsteady Chair", "Chess", 279},
	{"Bling", "Chess", 279},
	{"King", "Chess", 279},
}


func main() {
	fmt.Println("Hello formatting strings")
	fmt.Println("Product:", Kayak.Name, "Price:", Kayak.Price)
}