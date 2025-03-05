package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID   string  `gorm:"type:varchar(100)"`
	ProductID uint32  `gorm:"type:int(11)"`
	Quantity  uint32  `gorm:"type:int(11)"`
	Cost      float32 `gorm:"type:decimal(10,2)"`
}
