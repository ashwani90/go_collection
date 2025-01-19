package store

type Product struct {
	// Exported fields
	Name, Category string
	// Unexported Fields
	price float64
}