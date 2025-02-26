package service

import (
	"context"
	"errors"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/dal/mysql/model"
	_const "github.com/All-Done-Right/douyin-mall-microservice/app/order/const"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/global"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/order"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewRenewTokenByRPCService new RenewTokenByRPCService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}
func (s PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	if req.Address == nil {
		logrus.Errorln("address is empty")
		return nil, errors.New("address is empty")
	}
	orderID := uuid.NewString()
	orderData := model.Order{
		OrderID:      orderID,
		UserID:       req.UserId,
		UserCurrency: req.UserCurrency,
		Address: model.Address{
			Email:         req.Email,
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
		},
	}
	if len(req.OrderItems) == 0 {
		logrus.Errorln("order items is empty")
		return nil, errors.New("order items is empty")
	}
	for _, item := range req.OrderItems {
		orderData.OrderItems = append(orderData.OrderItems, model.OrderItem{
			ProductID: item.Item.ProductId,
			Quantity:  item.Item.Quantity,
			Cost:      item.Cost,
		})
	}
	err = global.RDB.Set(s.ctx, orderID, req.UserId, time.Second*_const.ORDER_TIME_TO_EXPIRE).Err()
	if err != nil {
		logrus.Errorln("set order id to redis failed")
		return nil, errors.New("set order id to redis failed")
	}
	if err := global.DB.Create(&orderData).Error; err != nil {
		logrus.Errorln("create order failed")
		return nil, errors.New("create order failed")
	}

	resp = &order.PlaceOrderResp{
		Order: &order.OrderResult{
			OrderId: orderID,
		},
	}
	return resp, nil

}
