package product

import (
	"time"
)

type ProductFormatter struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	UnitPrice   int       `json:"unitPrice"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Code        string    `json:"code"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FormatProductV1(product Product) ProductFormatter {
	productFormatter := ProductFormatter{}
	productFormatter.ID = product.ID
	productFormatter.Title = product.Title
	productFormatter.UnitPrice = product.UnitPrice
	productFormatter.Description = product.Description
	productFormatter.Image = product.Image
	productFormatter.Code = product.Code

	return productFormatter
}

func FormatProducts(products []Product) []ProductFormatter {
	productsFormatter := []ProductFormatter{}

	for _, product := range products {
		productFormatter := FormatProductV1(product)
		productsFormatter = append(productsFormatter, productFormatter)
	}

	return productsFormatter
}
