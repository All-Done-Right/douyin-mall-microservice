package service

import (
	"GoMall/app/product/biz/dal/mysql"
	"GoMall/app/product/biz/model"
	product "GoMall/rpc_gen/kitex_gen/product"
	"context"
	"fmt"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	fmt.Println("到product服务的搜索方法")
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	products, err := productQuery.SearchProducts(req.Query)
	var results []*product.Product
	for _, v := range products {
		results = append(results, &product.Product{
			Id:          uint32(v.ID),
			Name:        v.Name,
			Price:       v.Price,
			Picture:     v.Picture,
			Description: v.Description,
		})

	}
	fmt.Println("到product服务的搜索方法搜索结果")
	fmt.Println(&product.SearchProductsResp{Results: results})
	return &product.SearchProductsResp{Results: results}, err
}
