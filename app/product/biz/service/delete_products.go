package service

import (
	"context"
	"fmt"

	"github.com/All-Done-Right/douyin-mall-microservice/product/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/product/biz/model"
	product "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type DeleteProductsService struct {
	ctx context.Context
} // NewDeleteProductsService new DeleteProductsService
func NewDeleteProductsService(ctx context.Context) *DeleteProductsService {
	return &DeleteProductsService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductsService) Run(req *product.DeleteProductsReq) (resp *product.Empty, err error) {
	// Finish your business logic.
	// Finish your business logic.
	fmt.Println("到product服务的删除方法")
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id is required")
	}
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	err = productQuery.DeleteProductQuery(req.Id)
	if err != nil {
		return nil, err
	}
	fmt.Println("product服务的删除方法结束")
	return
}
