package service

import (
	"context"
	"errors"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/dal/repo/model"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/dal/repo/repo_dao"
	_const "github.com/All-Done-Right/douyin-mall-microservice/app/order/const"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/global"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/order"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
)

type PlaceOrderService struct {
	ctx context.Context
	db  repo_dao.Repo
} // NewRenewTokenByRPCService new RenewTokenByRPCService
func NewPlaceOrderService(ctx context.Context, db repo_dao.Repo) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx, db: db}

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
		Email:        req.Email,
		Address: model.Address{

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
		logrus.Errorln("set order id to cache failed")
		return nil, err
	}
	orderID, err = s.db.CreateOrder(orderData)
	if err != nil {
		logrus.Errorln("create order failed")
		return nil, err
	}

	resp = &order.PlaceOrderResp{
		Order: &order.OrderResult{
			OrderId: orderID,
		},
	}
	return resp, nil

}
