package organization

type OrganizationCreateInput struct {
	Name   string `json:"name" validate:"required"`
	Status int8   `json:"status" validate:"required"`
}

type FindById struct {
	ID int `uri:"id" binding:"required"`
}
