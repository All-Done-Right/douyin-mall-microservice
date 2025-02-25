package model

import (
	"context"
	"gorm.io/gorm"
)

type Category struct {
	Base
	CategoryName string    `json:"category_name"`
	Description  string    `json:"description"`
	Products     []Product `json:"products" gorm:"many2many:product_category"`
}

func (c Category) TableName() string {
	return "category"
}

type CategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (c CategoryQuery) GetProductsByCategoryName(name string) (categories []Category, err error) {
	err = c.db.WithContext(c.ctx).Model(&Category{}).Where(&Category{CategoryName: name}).Preload("Products").Find(&categories).Error
	return
}

func NewCategoryQuery(ctx context.Context, db *gorm.DB) *CategoryQuery {
	return &CategoryQuery{
		ctx: ctx,
		db:  db,
	}

}
