package service

import (
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/dal/repo/model"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepo struct {
	mock.Mock
}

func (m *MockOrderRepo) CreateOrder(order model.Order) (string, error) {
	args := m.Called(order)
	return args.String(0), args.Error(1)
}
func (m *MockOrderRepo) ListOrders(UserID uint32) ([]model.Order, error) {
	args := m.Called(UserID)
	// 第一个参数是返回的订单列表，第二个参数是错误
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Order), args.Error(1)
}
func (m *MockOrderRepo) MarkOrderPaid(orderID string) error {
	args := m.Called(orderID)

	return args.Error(0)
}
