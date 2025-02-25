package service

import (
	"GoMall/app/frontend/infra/rpc"
	rpcproduct "GoMall/rpc_gen/kitex_gen/product"
	"context"
	"fmt"

	product "GoMall/app/frontend/hertz_gen/frontend/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateProductService(Context context.Context, RequestContext *app.RequestContext) *UpdateProductService {
	return &UpdateProductService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateProductService) Run(req *product.UpdateProductsReq) (resp *product.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	fmt.Println("开始RPC调用Product服务")
	_, err = rpc.ProductClient.InsertProducts(h.Context, &rpcproduct.InsertProductsReq{Product: &rpcproduct.Product{
		Id:          req.Product.Id,
		Name:        req.Product.Name,
		Price:       req.Product.Price,
		Picture:     req.Product.Picture,
		Description: req.Product.Description,
	}})
	if err != nil {
		return nil, err
	}
	return
}
