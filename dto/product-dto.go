package dto

type ProductUpdateDTO struct {
	ID               uint64 `json:"id" form:"id" binding:"required"`
	Title            string `json:"title" form:"title" binding:"required"`
	ShortDescription string `json:"short_description" form:"short_description" binding:"required"`
	Description      string `json:"description" form:"description" binding:"required"`
	FreeDelivey      bool   `json:"free_delivery" form:"free_delivery" binding:"required"`
	Active           bool   `json:"active" form:"active" binding:"required"`
	Price            int    `json:"price" form:"price" binding:"required"`
	AdminID          uint64 `json:"admin_id,omitempty" form:"admin_id,omitempty"`
}

type ProductCreateDTO struct {
	Title            string `json:"title" form:"title" binding:"required"`
	ShortDescription string `json:"short_description" form:"short_description" binding:"required"`
	Description      string `json:"description" form:"description" binding:"required"`
	FreeDelivey      bool   `json:"free_delivery" form:"free_delivery" binding:"required"`
	Active           bool   `json:"active" form:"active" binding:"required"`
	Price            int    `json:"price" form:"price" binding:"required"`
	AdminID          uint64 `json:"admin_id,omitempty" form:"admin_id,omitempty"`
}

