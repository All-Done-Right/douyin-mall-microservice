package service

import (
	"context"
	"errors"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/dal/repo/repo_dao"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/global"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/order"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type MarkOrderPaidService struct {
	ctx context.Context
	db  repo_dao.Repo
} // NewRenewTokenByRPCService new RenewTokenByRPCService
func NewMarkOrderPaidService(ctx context.Context, db repo_dao.Repo) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx, db: db}
}
func (s MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	userID := req.GetUserId()
	orderID := req.GetOrderId()
	DBUserID, err := global.RDB.Get(s.ctx, orderID).Uint64()
	if errors.Is(err, redis.Nil) {
		logrus.Debugln("order not found,maybe is expired")
		return nil, errors.New("order not found")
	}
	if DBUserID != uint64(userID) {
		logrus.Debugln("order not belong to the user")
		return nil, errors.New("order not belong to the user")
	}
	err = s.db.MarkOrderPaid(orderID)
	if err != nil {
		logrus.Errorln(err)
		return nil, errors.New("mark order paid failed")
	}
	return resp, nil

}
