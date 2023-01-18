package order

import (
	"time"
	"tokokecilkita-go/product"
)

type Order struct {
	ID           int    `gorm:"size:36;not null;uniqueIndex;primary"`
	OrderNumber  string `gorm:size:32`
	Status       string `gorm:size:100`
	GrandTotal   int    `gorm:size:100`
	CustomerName string `gorm:size:100`
	UserID       int    `gorm:size:32`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type OrderProduct struct {
	ID        int `gorm:"size:36;not null;uniqueIndex;primary"`
	Qty       int `gorm:size:100`
	SubTotal  int `gorm:size:100`
	UnitPrice int `gorm:size:100`
	OrderID   int `gorm:size:100`
	ProductID int `gorm:size:100`
	Product   product.Product
	Order     Order
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderProductCreate struct {
	ID        int `gorm:"size:36;not null;uniqueIndex;primary"`
	Qty       int `gorm:size:100`
	SubTotal  int `gorm:size:100`
	UnitPrice int `gorm:size:100`
	OrderID   int `gorm:size:100`
	ProductID int `gorm:size:100`
	CreatedAt time.Time
	UpdatedAt time.Time
}
