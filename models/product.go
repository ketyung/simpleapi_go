package models

type Product struct {
	Code string `json:"code"`

	Name string `json:"name"`

	Qty int `json:"qty"`

	LastUpdated string `json:"last_updated"`
}

/**
In GO exported variable and function name
must start with big cap
*/
func SampleProducts() []Product {

	var products = []Product{

		{Code: "P0122", Name: "Test Product 111", Qty: 122, LastUpdated: "2021-12-19 18:18:29"},
		{Code: "P0123", Name: "Test Product 122", Qty: 69, LastUpdated: "2021-12-19 08:02:12"},
		{Code: "P0124", Name: "Test Product 133", Qty: 280, LastUpdated: "2021-12-19 13:30:30"},
	}

	return products
}
