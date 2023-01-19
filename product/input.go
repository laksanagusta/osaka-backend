package product

type ProductCreateInput struct {
	Title       string `json:"title" validate:"required"`
	UnitPrice   int    `json:"unitPrice" validate:"required"`
	Description string `json:"description" validate:"required"`
	Code        string `json:"code" validate:"required"`
}

type FindById struct {
	ID int `uri:"id" binding:"required"`
}

type FindByCode struct {
	Code string `uri:"code" binding:"required"`
}
