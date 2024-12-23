package main

import (
	"fmt"
	"math/rand"
	"sort"
)


// Basic Types values and pointers

func test1() {
	fmt.Println(rand.Int())
	fmt.Println("Value: ", rand.Int())
	fmt.Println(20+20)
	fmt.Println(20+30)
	const price float32 = 275.00
	const tax float32 = 27.50
	const quantity int = 2
	// mismatched type
	// fmt.Println("Total is : ", (price+tax)*quantity)
	const quantity2, inStock = 2, true
	// be more flexible with the type now
	fmt.Println("Total is : ", (price+tax)*quantity2)
	fmt.Println("In Stock: ", inStock)
}	

func VariableDecs() {
	var price float32 = 275.00
	var tax float32 = 27.50
	fmt.Println(price+tax)
	price = 300
	fmt.Println(price+tax)
	// We can also omit datatype
	var price2 = 275.00
	var price3 = price2
	fmt.Println(price2)
	fmt.Println(price3)

	// We can also omit variable's value assignment
	var price4 float32
	fmt.Println(price4)
	price4 = 275.00
	fmt.Println("Price 4: ", price4)

	// And other ways of declaring variables exist
	// Short variable declaration
	price7 := 280
	fmt.Println("Price 7: ", price7)

	// Blank identifier declaration
	var _ = "alice"
	fmt.Println("hello alice")
}

// Pointer usage
func PointerSimple() {
	first := 100
	second := first
	fmt.Println("First value ", first)
	fmt.Println("Second value ", second)
	// Now both first and second pointer to different address positions

	one := 100
	// Pointer to an int value
	var two *int = &one
	fmt.Println("First value ", one)
	fmt.Println("Second address ", two)
	fmt.Println("Second value ", *two)
	one++
	*two++

	var newPointer *int
	newPointer = two
	*newPointer++
	fmt.Println("First value ", one)
	fmt.Println("Second address ", two)
	fmt.Println("Second value ", *two)

	// Understanding pointer zero values
	pehla := 100
	var doosra *int

	fmt.Println(doosra)
	// this will give error because no pointer value exists
	// fmt.Println(*doosra)
	doosra = &pehla
	fmt.Println(doosra)
	fmt.Println(doosra==nil)

	uno := 100
	dos := &first
	tres := &dos

	// Following the pointer
	fmt.Println(uno)
	fmt.Println(*dos)
	fmt.Println(**tres)

}

func main() {

	names := [5]string{ "Hanuman", "Chandramukhi", "Gulgule", "Gopi", "Koki" }
	secondPosition2 := names[1]
	fmt.Println(secondPosition2)
	sort.Strings(names[:])
	fmt.Println(secondPosition2)

	names2 := [5]string{ "Hanuman", "Chandramukhi", "Gulgule", "Gopi", "Koki" }
	secondPosition := &names2[1]
	fmt.Println(*secondPosition)
	sort.Strings(names2[:])
	fmt.Println(*secondPosition)
}