package entity

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	ItemCode    string `json:"itemCode" binding:"required" gorm:"type:varchar(32)" example:"123"`
	Description string `json:"description" binding:"max=255" gorm:"type:varchar(255)" example:"IPhone 10X"`
	Quantity    int    `json:"quantity" binding:"required" example:"1"`
	OrderID     uint
}
