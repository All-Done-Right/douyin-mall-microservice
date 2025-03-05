package service

import (
	"context"
	"fmt"
	"github.com/All-Done-Right/douyin-mall-microservice/product/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/product/biz/model"
	product1 "github.com/All-Done-Right/douyin-mall-microservice/product/biz/model"
	product "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type InsertProductsService struct {
	ctx context.Context
} // NewInsertProductsService new InsertProductsService
func NewInsertProductsService(ctx context.Context) *InsertProductsService {
	return &InsertProductsService{ctx: ctx}
}

// Run create note info
func (s *InsertProductsService) Run(req *product.InsertProductsReq) (resp *product.Empty, err error) {
	// Finish your business logic.
	fmt.Println("到product服务的新增方法")
	if req.Product == nil {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product  is required")
	}
	if req.Product.Name == "" {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product Name is required")
	}
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	var p product1.Product
	p.Name = req.Product.Name
	p.Description = req.Product.Description
	p.ID = int(req.Product.Id)
	p.Picture = req.Product.Picture
	p.Price = req.Product.Price
	err = productQuery.InsertProductQuery(p, req.CategoryId)
	if err != nil {
		return nil, err
	}
	fmt.Println("product服务的新增方法结束")
	return
}
