package main

import (
	"fmt"
	"sort"
	"reflect"
	"strconv"
)

func main2() {
	// arrays in go
	var names [3]string
	names[0] = "Kayak"
	names[1] = "lifeJacket"
	names[2] = "Paddle"

	fmt.Println(names[1])

	// using array literanl syntax
	names = [3]string{"kayab", "abdj", "gfdgdf"}
	fmt.Println(names[2])
	fmt.Println(names)

	// usage of pointers
	otherNames := &names
	names[0] = "Canoe"
	fmt.Println(names)
	fmt.Println(*otherNames)

	for index, name := range names {
		fmt.Println("Index: ", index, " Name: ", name)
	}

	// Working with slices
	name2 := make([]string, 3)
	name2[0] = "hello"
	name2[1] = "hi"
	name2[2] = "ok"

	fmt.Println(name2)

	// using literal syntax
	name3 := []string {"abcd", "edge", "fefd"}
	fmt.Println(name3)

	// adding elements to the slice
	name3 = append(name3, "Hello", "Hi")
	fmt.Println(name3)

	// both variables have different places in memory
	// so usage of pointer is advised

	name4 := make([]string, 3, 6)
	name4[0] = "sdhu"
	name4[1] = "sdfsdhu"
	name4[2] = "sdhdfsu"

	fmt.Println("length is: ", len(name4))
	fmt.Println("capacity is: ", cap(name4))

	moreNames := []string { "ok another" }
	appendedNames := append(name4, moreNames...)
	fmt.Println("Appended name: ", appendedNames)
}

func main() {
	// Using slices from existing arrays
	products := [4]string {"One", "Two", "Three", "Four"}
	someProducts := products[1:3]
	allProducts := products[:]

	fmt.Println(someProducts)
	fmt.Println(allProducts)

	someProducts = append(someProducts, "Hello")
	fmt.Println(someProducts)

	ourNames := products[1:]
	ourExNames := make([]string, 2)
	copy(ourExNames, ourNames)
	fmt.Println(ourNames)
	fmt.Println(ourExNames)

	allProd := products[1:]
	someProd := []string {"Bot", "Shoes"}
	copy(someProd[1:], allProd[2:3])
	fmt.Println(allProd)
	fmt.Println(someProd)

	// deleting slice elements
	deleted := append(products[:2], products[3:]...)
	fmt.Println(deleted)

	for index, value := range products[1:] {
		fmt.Println(index, value)
	}

	// sorting slices
	sort.Strings(someProd)
	for index, value := range someProd {
		fmt.Println(index, value)
	}

	// Comparing slices with Deep equal
	someP := someProd
	fmt.Println("Equal are: ", reflect.DeepEqual(someP, someProd))

	// Getting underlying array of a slice
	arrayPtr := (*[2]string)(someProd)
	array := *arrayPtr
	fmt.Println(array)


	// Using maps
	items := make(map[string]float64, 10)
	items["ab"] = 1.2
	items["cd"] = 4.3

	fmt.Println("Map size: ", len(items))
	fmt.Println("Value at ab is: ", items["ab"])

	// using map literal syntax
	items2 := map[string]float64 {
		"dr": 2.4,
		"fr": 34.4,
	}

	fmt.Println("Map size: ", len(items2))
	fmt.Println("Value at ab is: ", items2["ab"])

	// check for an item in array
	value, ok := items["hat"]
	if (ok) {
		fmt.Println("Stored value ", value)
	} else {
		fmt.Println("No stored such value")
	}

	// using concise statement
	if value, ok := items["ab"]; ok {
		fmt.Println("Stored value ", value)
	} else {
		fmt.Println("No stored such value")
	}

	// Removing items from a map
	delete(items, "ab")
	if value, ok := items["ab"]; ok {
		fmt.Println("Stored value ", value)
	} else {
		fmt.Println("No stored such value")
	}

	// Enumerating contents of a map
	for key, value := range items {
		fmt.Println("Key is: ", key, " Value is: ", value)
	}

	// Enumerating map in order
	keys := make([]string, 0, len(items2))
	for key, _ := range items2 {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("Key: ", key, " value: ", items2[key])
	}

	// Dual nature of strings
	var price string = "$24.45"
	var currency byte = price[0]
	var amountString string = price[1:]
	amount, parseErr := strconv.ParseFloat(amountString, 64)
	fmt.Println("Currency is: ", currency)
	// this is good when a character can be filled into a byte
	// but for larger size character we need rune
	if parseErr == nil {
		fmt.Println("Amount is: ", amount)
	} else {
		fmt.Println("Parse Error is: ", parseErr)
	}

	// Converting strings to runes
	var price2 []rune = []rune("$48.85")
	var currency2 string = string(price2[0])
	var amountString2 string = string(price2[1:])
	amount2, parsErr := strconv.ParseFloat(amountString2, 64)
	fmt.Println("Currency is: ", currency2)
	// this is good when a character can be filled into a byte
	// but for larger size character we need rune
	if parsErr == nil {
		fmt.Println("Amount is: ", amount2)
	} else {
		fmt.Println("Parse Error is: ", parsErr)
	}

	// Enumerating strings
	for index, character := range price2 {
		fmt.Println(index, character, string(character))
	}

	for index, character := range []byte(price) {
		fmt.Println(index, character)
	}
}