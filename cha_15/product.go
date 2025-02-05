package main

import "strconv"

type Product struct {
	Name, Category string
	Price float64
}

type ProductSlice []*Product


var Products = ProductSlice {
	{"Kayak", "Watersports", 279},
	{ "lifeJacket", "Watersports", 300 },
	{ "lifeJacket2", "Watersports", 360 },
	{ "lifeJacket3", "Watersports", 350 },
	{ "lifeJacke4", "Watersports", 340 },
	{ "lifeJacket4", "Watersports", 330 },
	{ "lifeJacket5", "Watersports", 320 },
	{ "lifeJacket6", "Watersports", 310 },
}

func ToCurrency(val float64) string {
	return "$" + strconv.FormatFloat(val, 'f', 2, 64)
}

