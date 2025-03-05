package service

import (
	"context"
	"fmt"
	"github.com/All-Done-Right/douyin-mall-microservice/app/frontend/infra/rpc"
	rpcproduct "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"

	product "github.com/All-Done-Right/douyin-mall-microservice/app/frontend/hertz_gen/frontend/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	fmt.Println("开始RPC调用Product服务")
	products, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{
		Query: req.Q,
	})
	if err != nil {

		return nil, err
	}
	if products == nil {
		fmt.Println("RPC调用User服务失败search")
		fmt.Println("products == nil")
		fmt.Println("products")
	}

	return utils.H{
		"items": products.Results,
		"q":     req.Q}, nil
}
