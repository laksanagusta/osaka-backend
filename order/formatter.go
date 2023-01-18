package order

import (
	"time"
	"tokokecilkita-go/product"
)

type OrderFormatter struct {
	ID           int                       `json:"id"`
	Status       string                    `json:"status"`
	GrandTotal   int                       `json:"grandTotal"`
	CustomerName string                    `json:"customerName"`
	OrderNumber  string                    `json:"orderNumber"`
	OrderProduct []OrderProductFormatterV2 `json:"orderProducts"`
	CreatedAt    time.Time                 `json:"created_at"`
	UpdatedAt    time.Time                 `json:"updated_at"`
}

type OrderProductFormatter struct {
	ID        int                      `json:"id"`
	Qty       int                      `json:"qty"`
	SubTotal  int                      `json:"subTotal"`
	UnitPrice int                      `json:"unitPrice"`
	OrderID   int                      `json:"orderId"`
	ProductID int                      `json:"productId"`
	Product   product.ProductFormatter `json:"product"`
	Order     OrderFormatter           `json:"order"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedAt time.Time                `json:"updated_at"`
}

type OrderProductFormatterV2 struct {
	ID        int       `json:"id"`
	Qty       int       `json:"qty"`
	SubTotal  int       `json:"subTotal"`
	UnitPrice int       `json:"unitPrice"`
	ProductID int       `json:"productId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatOrderV1(order Order) OrderFormatter {
	orderFormatter := OrderFormatter{}
	orderFormatter.ID = order.ID
	orderFormatter.Status = order.Status
	orderFormatter.GrandTotal = order.GrandTotal
	orderFormatter.OrderNumber = order.OrderNumber
	orderFormatter.CustomerName = order.CustomerName

	return orderFormatter
}

func FormatOrders(orders []Order) []OrderFormatter {
	ordersFormatter := []OrderFormatter{}

	for _, order := range orders {
		orderFormatter := FormatOrderV1(order)
		ordersFormatter = append(ordersFormatter, orderFormatter)
	}

	return ordersFormatter
}

func FormatOrderProduct(orderProduct OrderProduct) OrderProductFormatter {
	orderProductFormatter := OrderProductFormatter{}

	orderProductFormatter.ID = orderProduct.ID
	orderProductFormatter.Qty = orderProduct.Qty
	orderProductFormatter.SubTotal = orderProduct.SubTotal
	orderProductFormatter.UnitPrice = orderProduct.UnitPrice

	orderFormatter := OrderFormatter{}
	orderFormatter.ID = orderProduct.Order.ID
	orderFormatter.Status = orderProduct.Order.Status
	orderFormatter.GrandTotal = orderProduct.Order.GrandTotal
	orderFormatter.OrderNumber = orderProduct.Order.OrderNumber

	orderProductFormatter.Order = orderFormatter

	productFormatter := product.ProductFormatter{}
	productFormatter.ID = orderProduct.Product.ID
	productFormatter.Title = orderProduct.Product.Title
	productFormatter.Description = orderProduct.Product.Description
	productFormatter.UnitPrice = orderProduct.Product.UnitPrice
	productFormatter.Image = orderProduct.Product.Image

	orderProductFormatter.Product = productFormatter

	return orderProductFormatter
}

func FormatOrderProducts(orderProduct []OrderProduct) []OrderProductFormatter {
	orderProductsFormatter := []OrderProductFormatter{}

	for _, orderProduct := range orderProduct {
		orderProductFormatter := FormatOrderProduct(orderProduct)
		orderProductsFormatter = append(orderProductsFormatter, orderProductFormatter)
	}

	return orderProductsFormatter
}

func FormatOrderProductBasketV1(orderProduct OrderProduct) OrderProductFormatterV2 {
	orderProductFormatterV2 := OrderProductFormatterV2{}

	orderProductFormatterV2.ID = orderProduct.ID
	orderProductFormatterV2.Qty = orderProduct.Qty
	orderProductFormatterV2.SubTotal = orderProduct.SubTotal
	orderProductFormatterV2.UnitPrice = orderProduct.UnitPrice
	orderProductFormatterV2.ProductID = orderProduct.ProductID

	return orderProductFormatterV2
}

func FormatOrderBasketV1(order Order, orderProduct []OrderProduct) OrderFormatter {
	orderFormatter := OrderFormatter{}
	orderFormatter.ID = order.ID
	orderFormatter.Status = order.Status
	orderFormatter.GrandTotal = order.GrandTotal
	orderFormatter.CustomerName = order.CustomerName

	for _, orderProduct := range orderProduct {
		orderProductFormatter := FormatOrderProductBasketV1(orderProduct)
		orderFormatter.OrderProduct = append(orderFormatter.OrderProduct, orderProductFormatter)
	}

	return orderFormatter
}
