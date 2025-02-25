package service

import (
	"GoMall/app/product/biz/dal/mysql"
	"GoMall/app/product/biz/model"
	product "GoMall/rpc_gen/kitex_gen/product"
	"context"
	"fmt"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	fmt.Println("到product服务的搜索列表方法")
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)
	c, err := categoryQuery.GetProductsByCategoryName(req.CategoryName)
	if err != nil {
		return nil, err
	}
	resp = &product.ListProductsResp{}
	for _, v1 := range c {
		for _, v := range v1.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(v.ID),
				Name:        v.Name,
				Price:       v.Price,
				Description: v.Description,
				Picture:     v.Picture,
			})
		}
	}
	fmt.Println("到product服务的搜索列表方法搜索结果")
	fmt.Println(resp)
	return resp, nil
}
