package product

type ProductCreateInput struct {
	Title       string `json:"title" validate:"required" binding:"required"`
	UnitPrice   int    `json:"unitPrice" validate:"required" binding:"required"`
	Description string `json:"description" validate:"required" binding:"required"`
	Code        string `json:"code" validate:"required" binding:"required"`
}

type FindById struct {
	ID int `uri:"id" binding:"required"`
}

type FindByCode struct {
	Code string `uri:"code" binding:"required"`
}
