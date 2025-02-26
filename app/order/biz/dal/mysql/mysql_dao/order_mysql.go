package mysql_dao

import (
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/dal/mysql/model"
	"gorm.io/gorm"
)

type OrderRepo struct {
	*gorm.DB
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{
		db,
	}
}
func (r *OrderRepo) CreateOrder(order model.Order) (string, error) {
	if r.Create(&order).Error != nil {
		return "", r.Error
	}
	return order.OrderID, nil
}
func (r *OrderRepo) ListOrders(UserID uint32) ([]*model.Order, error) {

	var orders []*model.Order
	if err := r.Model(&model.Order{}).Where("user_id = ?", UserID).Preload("OrderItems").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
func (r *OrderRepo) MarkOrderPaid(OrderID string) error {
	return r.Where("order_id = ?", OrderID).Update("paid", true).Error
}
