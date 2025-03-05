package main

import (
	"context"

	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/model"
	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/service"
	cart "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/cart"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product/productcatalogservice"
)

// CartServiceImpl 实现购物车服务接口
type CartServiceImpl struct {
	store         model.CartStore
	productClient productcatalogservice.Client // 新增字段
}

func NewCartServiceImpl(
	store model.CartStore,
	productClient productcatalogservice.Client, // 新增参数
) *CartServiceImpl {
	return &CartServiceImpl{
		store:         store,
		productClient: productClient, // 初始化
	}
}

func (s *CartServiceImpl) AddItem(ctx context.Context, req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// 传递 s.productClient
	addItemService := service.NewAddItemService(ctx, s.store, s.productClient)
	resp, err = addItemService.Run(req)
	return resp, err
}

// GetCart 实现 CartService 接口的 GetCart 方法
func (s *CartServiceImpl) GetCart(ctx context.Context, req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// 调用 service.NewGetCartService 时提供正确的参数
	getCartService := service.NewGetCartService(ctx, s.store)
	resp, err = getCartService.Run(req)
	return resp, err
}

// EmptyCart 实现 CartService 接口的 EmptyCart 方法
func (s *CartServiceImpl) EmptyCart(ctx context.Context, req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	// 调用 service.NewEmptyCartService 时提供正确的参数
	emptyCartService := service.NewEmptyCartService(ctx, s.store)
	resp, err = emptyCartService.Run(req)
	return resp, err
}

