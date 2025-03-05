package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `gorm:"unique_index"`
	PasswordHashed string `gorm:"type:varchar(255) not null"`
	Nickname       string `gorm:"type:varchar(50)"`
	Avatar         string `gorm:"type:varchar(255)"`
	Phone          string `gorm:"type:varchar(20)"`
	Address        string `gorm:"type:varchar(255)"`
}

func (User) TableName() string {
	return "user"
}

func GetByEmail(db *gorm.DB, ctx context.Context, email string) (*User, error) {
	var user User
	err := db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 记录不存在时返回 nil, nil
		}
		return nil, err // 其他错误返回 nil, err
	}
	return &user, nil // 找到记录时返回 &user, nil
}

func Create(db *gorm.DB, ctx context.Context, user *User) error {
	return db.WithContext(ctx).Create(user).Error
}

func GetByID(db *gorm.DB, ctx context.Context, id uint) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where("id = ?", id).First(&user).Error
	return
}

func Update(db *gorm.DB, ctx context.Context, user *User) error {
	return db.WithContext(ctx).Save(user).Error
}
