package dto

type RegisterUserDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Password string `json:"password" form:"password" validate:"min:8" binding:"required"`
}
