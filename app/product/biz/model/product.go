package model

import (
	"context"
	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Picture     string     `json:"picture"`
	Price       float32    `json:"price"`
	Categorys   []Category `json:"categorys" gorm:"many2many:product_category"`
}

func (p Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (p ProductQuery) GetById(productId int) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, productId).Error
	return
}
func (p ProductQuery) SearchProducts(q string) (products []*Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&products, "name like ? or description like ?", "%"+q+"%", "%"+q+"%").Error
	return
}
func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db:  db,
	}

}

// 添加
func (p ProductQuery) InsertProductQuery(product Product, categoryId int32) (err error) {
	//插入数据
	err = p.db.WithContext(p.ctx).Create(product).Error
	//获取自增Id
	var id []int
	p.db.Raw("select LAST_INSERT_ID() as id").Pluck("id", &id)
	//新增分类_商品
	sql := "INSERT INTO product_category (category_id, product_id) VALUES( ? ,?)"
	p.db.WithContext(p.ctx).Exec(sql, categoryId, id[0])
	return
}

// 修改
func (p ProductQuery) UpdateProductQuery(product Product) (err error) {
	err = p.db.WithContext(p.ctx).Updates(product).Error
	return
}

// 删除
func (p ProductQuery) DeleteProductQuery(productId uint32) (err error) {
	err = p.db.WithContext(p.ctx).Delete(productId).Error
	//删除分类_商品中间表的数据
	sql := "DELETE FROM product_category where product_id = ?"
	err = p.db.WithContext(p.ctx).Exec(sql, productId).Error
	return
}
