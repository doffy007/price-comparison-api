package dto

type LoginAdminDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Market   string `json:"market" form:"market" binding:"required"`
	Password string `json:"password" form:"password" validate:"min:8" binding:"required"`
}