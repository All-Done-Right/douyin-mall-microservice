package service

import (
	"context"
	"fmt"
	"github.com/All-Done-Right/douyin-mall-microservice/app/frontend/infra/rpc"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"

	category "github.com/All-Done-Right/douyin-mall-microservice/app/frontend/hertz_gen/frontend/category"
	"github.com/cloudwego/hertz/pkg/app"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	fmt.Println("开始RPC调用Product服务")
	p, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{CategoryName: req.Category})
	if err != nil {
		return nil, err
	}

	if p == nil {
		fmt.Println("RPC调用Product服务失败Category")
		fmt.Println("p == nil")
		fmt.Println("p")
	}

	return utils.H{"title": "Category",
		"items": p.Products}, nil
}
