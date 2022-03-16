package dto

type AdminUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Market   string `json:"market" form:"market" binding:"required"`
	Password string `json:"password" form:"password" validate:"min:8" binding:"required"`
}

type AdminCreateDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Market   string `json:"market" form:"market" binding:"required"`
	Password string `json:"password" form:"password" validate:"min:8" binding:"required"`
}