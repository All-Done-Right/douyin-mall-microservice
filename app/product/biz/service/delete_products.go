package service

import (
	"GoMall/app/product/biz/dal/mysql"
	"GoMall/app/product/biz/model"
	product "GoMall/rpc_gen/kitex_gen/product"
	"context"
	"fmt"
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
