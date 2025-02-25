package model

import "time"

type Base struct {
	ID        int `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
