package service

import (
	"context"
	"fmt"
	"github.com/All-Done-Right/douyin-mall-microservice/app/frontend/hertz_gen/frontend/product"
	"github.com/All-Done-Right/douyin-mall-microservice/app/frontend/infra/rpc"
	rpcproduct "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	fmt.Println("开始RPC调用Product服务")
	p, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{
		Id: req.Id,
	})
	if p == nil {
		fmt.Println("RPC调用Product服务失败get_product")
		fmt.Println("p == nil")
		fmt.Println("p")
	}

	if err != nil {
		return nil, err
	}
	return utils.H{"items": p.Product}, nil
}
