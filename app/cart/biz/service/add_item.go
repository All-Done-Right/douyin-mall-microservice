package service

import (
	"context"

	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/model"
	cart "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/cart"
	product "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product"
	"gorm.io/gorm"

	RPCproduct "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type AddItemService struct {
	ProductClient RPCproduct.Client
	CartStore     model.CartStore
	Ctx           context.Context
	DB            *gorm.DB
}

// NewAddItemService new AddItemService
func NewAddItemService(
	ctx context.Context,
	store model.CartStore,
	productClient RPCproduct.Client,
) *AddItemService {
	return &AddItemService{
		Ctx:           ctx,
		CartStore:     store,
		ProductClient: productClient,
	}
}

// Run create note info

func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	if s.ProductClient == nil {
		return nil, kerrors.NewBizStatusError(50000, "product client is nil")
	}
	if s.CartStore == nil {
		return nil, kerrors.NewBizStatusError(50000, "cart store is nil")
	}
	if s.ProductClient == nil {
		panic("productClient is nil")
	}
	//productResp, err := rpc.ProductClient.GetProduct(s.Ctx, &product.GetProductReq{Id: req.Item.ProductId})
	productResp, err := s.ProductClient.GetProduct(s.Ctx, &product.GetProductReq{Id: req.Item.ProductId})

	if err != nil {
		return nil, err
	}
	if productResp == nil || productResp.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(40004, "product not found")
	}
	cartItem := &model.Cart{
		UserID:    req.UserId,
		ProductID: req.Item.ProductId,
		Qty:       req.Item.Quantity,
	}
	err = s.CartStore.AddItem(s.Ctx, cartItem)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}
	return &cart.AddItemResp{}, nil
}
