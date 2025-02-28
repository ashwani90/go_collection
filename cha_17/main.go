package main


import "fmt"

// type Product struct {
// 	Name, Category string
// 	Price float64
// }

// var Kayak = Product{
// 	Name: "Kayak",
// 	Category: "Watersports",
// 	Price: 275,
// }

// var Products =[]Product {
// 	{"Keyak", "watersports", 279},
// 	{"Lifejacket", "watersports", 279},
// 	{"Soccer Ball", "Soccer", 279},
// 	{"Corner Flags", "Soccer", 279},
// 	{"Stadium", "Soccer", 279},
// 	{"Thinking Cap", "Chess", 279},
// 	{"Unsteady Chair", "Chess", 279},
// 	{"Bling", "Chess", 279},
// 	{"King", "Chess", 279},
// }


func getProductName(index int) (name string,  err error) {
	if (len(Products) > index) {
		name = fmt.Sprintf("Name of product: %v", Products[index].Name)
	} else {
		err = fmt.Errorf("Error for index %v", index)
	}
	return
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template + "\n", values...)
}

func (p Product) String() string {
	return fmt.Sprintf("Product: %v, Price: $%4.2f", p.Name, p.Price)
}

func main() {
	fmt.Println("Hello formatting strings")
	fmt.Println("Product:", Kayak.Name, "Price:", Kayak.Price)

	fmt.Printf("Product: %v, Price: $%4.2f", Kayak.Name, Kayak.Price)

	name, _ := getProductName(1)
	fmt.Println(name)

	_, err := getProductName(10)
	fmt.Println(err.Error())

	Printfln("Value %v", Kayak)
	Printfln("Go syntax: %#v", Kayak)
	Printfln("Type: %T", Kayak)

	Printfln("Value with fields: %+v", Kayak)

	number := 250
	Printfln("Binary: %b", number)
	Printfln("Decimal: %d", number)
	Printfln("Octal: %o, %O", number)
	Printfln("Hexadecimal: %x, %X", number)
}