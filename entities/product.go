package entities

//represent of database
type Product struct {
	ID               uint64 `gorm:"primaryKey:autoIncremet" json:"id"`
	Title            string `gorm:"type:varchar(255)" json:"title"`
	ShortDescription string `gorm:"type:varchar(255)" json:"type"`
	Description      string `gorm:"type:varchar(255)" json:"description"`
	FreeDelivey      bool   `gorm:"type:bool" json:"free_delivery"`
	Active           bool   `gorm:"type:bool" json:"status_active"`
	Price            int    `gorm:"type:int" json:"price"`
	AdminID          uint64 `gorm:"not null" json:"-"`
	Admin            Admin  `gorm:"foreignKey:AdminID;constraint:OnUpdate:OnDelete:CASCADE" json:"admin"`
}
