package main

import (
	"context"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/service"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/order"
	"github.com/sirupsen/logrus"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	srv := service.NewPlaceOrderService(ctx)
	resp, err = srv.Run(req)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	return resp, nil
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	srv := service.NewListOrderService(ctx)
	resp, err = srv.Run(req)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	return resp, nil
}

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	srv := service.NewMarkOrderPaidService(ctx)
	resp, err = srv.Run(req)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	return resp, nil
}
