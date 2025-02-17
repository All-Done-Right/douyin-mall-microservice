package model

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `gorm:"unique_index"`
	PasswordHashed string `gorm:"type:varchar(255) not null"`
}

func (User) TableName() string {
	return "user"
}

func GetByEmail(db *gorm.DB, ctx context.Context, email string) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where(&User{Email: email}).First(&user).Error
	return
}

func Create(db *gorm.DB, ctx context.Context, user *User) error {
	return db.WithContext(ctx).Create(user).Error
}

func Delete(db *gorm.DB, ctx context.Context, id uint) error {
	return db.WithContext(ctx).Delete(&User{}, id).Error
}

func Update(db *gorm.DB, ctx context.Context, user *User) error {
	return db.WithContext(ctx).Save(user).Error
}

func GetByID(db *gorm.DB, ctx context.Context, id uint) (user *User, err error) {
	err = db.WithContext(ctx).First(&user, id).Error
	return
}
