package main

import "strconv"

type Product struct {
	Name, Category string
	Price float64
}

var ProductList = []*Product {
	{"Kayak", "Watersorts", 279},
	{"Kayak 2", "Watersorts 1", 2792},
	{"Kayak 3", "Watersorts 2", 2791},
}

type ProductGroup []*Product

type ProductData = map[string]ProductGroup

var Products = make(ProductData)

func ToCurrency(val float64) string {
	return "$" + strconv.FormatFloat(val, 'f', 2, 64)
}

func init() {
	for _,p := range ProductList {
		if _, ok := Products[p.Category];ok {
			Products[p.Category] = append(Products[p.Category], p)
		} else {
			Products[p.Category] = ProductGroup{p}
		}
	}
}