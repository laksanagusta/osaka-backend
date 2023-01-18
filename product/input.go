package product

type ProductCreateInput struct {
	Title       string `json:"title" validate:"required"`
	UnitPrice   int    `json:"unitPrice" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type FindById struct {
	ID int `uri:"id" binding:"required"`
}
