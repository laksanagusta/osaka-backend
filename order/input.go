package order

type OrderCreateInput struct {
	Status       string `json:"status" validate:"required"`
	GrandTotal   int    `json:"grandTotal"`
	CustomerName string `json:"customerName" validate:"required"`
	OrderNumber  string `json:"orderNumber" validate:"required"`
	UserID       int    `json:"userId" validate:"required"`
}

type FindById struct {
	ID int `uri:"id" binding:"required"`
}

type OrderCreateV2Input struct {
	Order  OrderCreateInput     `json:"order" validate:"required"`
	Basket []OrderProductCreate `json:"basket" validate:"required"`
}

type UpdateBasketInput struct {
	Basket []OrderProductCreate `json:"basket" validate:"required"`
}
