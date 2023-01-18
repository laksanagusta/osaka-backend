package user

type RegisterUserInput struct {
	Username   string `json:"username" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	Role       string `json:"role"`
}

type UpdateUserInput struct {
	Username   string `json:"username" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
}

type UpdateUserRoleInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Role       string `json:"role"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type GetUserDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
