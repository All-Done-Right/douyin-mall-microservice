package service

import (
	"context"
	"fmt"
	"github.com/All-Done-Right/douyin-mall-microservice/product/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/product/biz/model"
	product "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.

	fmt.Println("到product服务的查询方法")
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id is required")
	}
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	p, err := productQuery.GetById(int(req.Id))
	if err != nil {
		return nil, err
	}
	fmt.Println("到product服务的搜索方法")
	fmt.Println(p)
	return &product.GetProductResp{
		Product: &product.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Price:       p.Price,
			Picture:     p.Picture,
			Description: p.Description,
		},
	}, nil
}
