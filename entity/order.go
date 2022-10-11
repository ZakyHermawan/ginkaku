package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	gorm.Model
	CustomerName string    `json:"customerName" binding:"required" gorm:"type:varchar(100)" example:"Tom Jerry"`
	OrderedAt    time.Time `json:"orderedAt" gorm:"default:CURRENT_TIMESTAMP" example:"2019-11-09T21:21:46+00:00"`
	Items        []Item    `gorm:"foreignKey:OrderID"`
}
