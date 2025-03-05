package main

import (
	"context"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/dal/repo/repo_dao"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/service"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/global"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/order"
	"github.com/sirupsen/logrus"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	db := repo_dao.NewOrderRepo(global.DB)
	srv := service.NewPlaceOrderService(ctx, db)
	resp, err = srv.Run(req)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	return resp, nil
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	db := repo_dao.NewOrderRepo(global.DB)
	srv := service.NewListOrderService(ctx, db)
	resp, err = srv.Run(req)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	return resp, nil
}

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	db := repo_dao.NewOrderRepo(global.DB)
	srv := service.NewMarkOrderPaidService(ctx, db)
	resp, err = srv.Run(req)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	return resp, nil
}
