package model

import "time"

type Base struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdateAt  time.Time
}
