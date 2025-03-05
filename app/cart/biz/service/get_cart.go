package service

import (
	"context"

	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/model"
	"gorm.io/gorm"

	cart "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetCartService struct {
	CartStore model.CartStore
	DB        *gorm.DB // 新增DB字段
	Ctx       context.Context
}

// NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context, store model.CartStore) *GetCartService {
	return &GetCartService{
		Ctx:       ctx,
		CartStore: store,
	}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	//list, err := model.GetCartByUserId(mysql.DB, s.Ctx, req.UserId)
	list, err := s.CartStore.GetCartByUserId(s.Ctx, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50002, err.Error())
	}
	var items []*cart.CartItem
	for _, item := range list {
		items = append(items, &cart.CartItem{
			ProductId: item.ProductID,
			Quantity:  item.Qty,
		})
	}
	return &cart.GetCartResp{Items: items}, nil
}
