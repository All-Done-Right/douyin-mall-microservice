package service

import (
	"context"
	"errors"

	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/dal/repo/repo_dao"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/cart"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/order"
	"github.com/sirupsen/logrus"
)

type ListOrderService struct {
	db  repo_dao.Repo
	ctx context.Context
} // NewRenewTokenByRPCService new RenewTokenByRPCService
func NewListOrderService(ctx context.Context, db repo_dao.Repo) *ListOrderService {
	return &ListOrderService{ctx: ctx, db: db}
}
func (s ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	userID := req.GetUserId()
	orders, err := s.db.ListOrders(userID)
	if err != nil {
		logrus.Errorln("list order failed")
		return nil, errors.New("list order failed")
	}
	logrus.Debugln(orders)
	resp = &order.ListOrderResp{
		Orders: make([]*order.Order, 0),
	}
	for _, orderData := range orders {
		orderItems := make([]*order.OrderItem, 0)
		for _, item := range orderData.OrderItems {
			orderItems = append(orderItems, &order.OrderItem{
				Cost: item.Cost,
				Item: &cart.CartItem{
					ProductId: item.ProductID,
					Quantity:  item.Quantity,
				},
			})
		}
		resp.Orders = append(resp.Orders, &order.Order{
			OrderId:      orderData.OrderID,
			UserId:       orderData.UserID,
			UserCurrency: orderData.UserCurrency,
			Email:        orderData.Email,
			Address: &order.Address{
				StreetAddress: orderData.Address.StreetAddress,
				City:          orderData.Address.City,
				State:         orderData.Address.State,
				Country:       orderData.Address.Country,
				ZipCode:       orderData.Address.ZipCode,
			},
			OrderItems: orderItems,
		})

	}
	return resp, nil
}
