package repo_dao

import "github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/dal/repo/model"

type Repo interface {
	CreateOrder(order model.Order) (string, error)
	ListOrders(UserID uint32) ([]model.Order, error)
	MarkOrderPaid(OrderID string) error
}
