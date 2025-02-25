package service

import (
	product "GoMall/app/frontend/hertz_gen/frontend/product"
	"GoMall/app/frontend/infra/rpc"
	rpcproduct "GoMall/rpc_gen/kitex_gen/product"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteProductService(Context context.Context, RequestContext *app.RequestContext) *DeleteProductService {
	return &DeleteProductService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteProductService) Run(req *product.DeleteProductsReq) (resp *product.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	fmt.Println("开始RPC调用Product服务")
	_, err = rpc.ProductClient.DeleteProducts(h.Context, &rpcproduct.DeleteProductsReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return
}
