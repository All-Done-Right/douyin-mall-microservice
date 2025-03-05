package model

import "time"

/*
type Base struct {
	ID        int32     `gorm:"primarykey;autoIncrement"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
*/
//解决MYSQL版本过低，数据类型不兼容问题
type Base struct {
	ID        int32     `gorm:"primarykey;autoIncrement"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
