package store

type Product struct {
	// Exported fields
	Name, Category string
	// Unexported Fields
	price float64
}

func NewProduct(name, category string, price float64) *Product {
	return &Product{ name, category, price }
}

func (p *Product) Price() float64 {
	return p.price
}

func (p *Product) SetPrice(newPrice float64) {
	p.price = newPrice
}

