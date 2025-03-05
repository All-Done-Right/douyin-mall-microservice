package service

import (
	"context"

	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/model"

	cart "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type EmptyCartService struct {
	CartStore model.CartStore
	Ctx       context.Context
}

// NewEmptyCartService new EmptyCartService
func NewEmptyCartService(Ctx context.Context, store model.CartStore) *EmptyCartService {
	return &EmptyCartService{
		Ctx:       Ctx,
		CartStore: store,
	}
}

// Run create note info

func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	err = s.CartStore.EmptyCart(s.Ctx, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50001, err.Error())
	}
	return &cart.EmptyCartResp{}, nil
}
