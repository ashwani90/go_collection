package main

import (
	"math"
	"fmt"
	"strconv"
)

func main1() {
	fmt.Println("Hello Guys")

	price, tax := 271.5, 4.4
	total := price * tax
	fmt.Println(total)

	var intVal = math.MaxInt64
	var floatVal = math.MaxFloat64
	fmt.Println(intVal*2)
	fmt.Println(floatVal*2)
	fmt.Println(math.IsInf((floatVal*2), 0))

	remainderResult := 3%2
	print(remainderResult)

	// Pointer comparison
	first := 100
	second := &first
	third := &first

	alpha := 100
	beta := &alpha

	// This is address comparison
	fmt.Println(second==third)
	fmt.Println(second==beta)

	// But this is value comparison
	fmt.Println(*second==*third)
	fmt.Println(*second==*beta)

	// Logical operations
	maxMph := 50
	passengerCapacity := 4
	airBags := true

	familyCar := passengerCapacity > 2 && airBags
	sportsCar := maxMph > 100 || passengerCapacity == 2
	canCategorize := !familyCar && !sportsCar

	fmt.Println(familyCar)
	fmt.Println(sportsCar)
	fmt.Println(canCategorize)

	// Explicit type conversion
	kayak := 275
	soccerBall := 19.50
	total = float64(kayak) + soccerBall
	fmt.Println(total)

	// Parsing from strings
	val1 := "true"
	val2 := "false"
	val3 := "not true"
	bool1, b1err := strconv.ParseBool(val1)
	bool2, b2err := strconv.ParseBool(val2)
	bool3, b3err := strconv.ParseBool(val3)
	if (b3err == nil) {
		fmt.Println("Parsed value ", bool3)
	} else {
		fmt.Println("Value is ", val3)
	}

	fmt.Println("Value 1 ", bool1, b1err)
	fmt.Println("Value 2 ", bool2, b2err)
	fmt.Println("Value 3 ", bool3, b3err)

	// Parsing Integeres
	val10 := 100
	int10, int10err := strconv.ParseInt(val1, 0, 8)
	if int10err == nil {
		fmt.Println("Parsed Value ", int10)
	} else {
		fmt.Println("Cannot parse ", val10)
	}
}

// more parsing
func main() {
	val1 := "100"
	int1, interr := strconv.ParseInt(val1, 0, 8)
	if interr == nil {
		smallInt := int8(int1)
		fmt.Println("Parsed value ", smallInt)
	} else {
		fmt.Println("Cannot parse value ", val1, interr)
	}

	// two step process 
	val2 := "100"
	int3, int2err := strconv.ParseInt(val2, 10, 0)
	if int2err == nil {
		var intResult int = int(int3)
		fmt.Println("Parsed value ", intResult)
	} else {
		fmt.Println("Cannot parse value ", val2, int2err)
	}

	// In one step this looks like
	valn := "100"
	intn, intnerr := strconv.Atoi(valn)
	if intnerr == nil {
		var intResult int = intn
		fmt.Println("Parse value ", intResult)
	} else {
		fmt.Println("Cannot parse ", valn, intnerr)
	}
	// Same thing we can do with float values

	valb := true
	valc := false
	strb := strconv.FormatBool(valb)
	strc := strconv.FormatBool(valc)
	fmt.Println("Formatted value 1 " + strb)
	fmt.Println("Formatted value 2 " + strc)

	// Conversion from int to string
	vali := 275
	base10string := strconv.Itoa(vali)

	fmt.Println("Base 10: " + base10string)

	// formatting float values
	valf := 49.95
	Fstring := strconv.FormatFloat(valf, 'f', 2, 64)
	Estring := strconv.FormatFloat(valf, 'e', -1, 64)

	fmt.Println("Format F: " + Fstring)
	fmt.Println("Format E: " + Estring)
}

