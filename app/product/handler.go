package main

import (
	"context"
	"github.com/All-Done-Right/douyin-mall-microservice/product/biz/service"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	resp, err = service.NewListProductsService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	resp, err = service.NewSearchProductsService(ctx).Run(req)

	return resp, err
}

// InsertProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) InsertProducts(ctx context.Context, req *product.InsertProductsReq) (resp *product.Empty, err error) {
	resp, err = service.NewInsertProductsService(ctx).Run(req)

	return resp, err
}

// DeleteProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) DeleteProducts(ctx context.Context, req *product.DeleteProductsReq) (resp *product.Empty, err error) {
	resp, err = service.NewDeleteProductsService(ctx).Run(req)

	return resp, err
}

// UpdateProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) UpdateProducts(ctx context.Context, req *product.UpdateProductsReq) (resp *product.Empty, err error) {
	resp, err = service.NewUpdateProductsService(ctx).Run(req)

	return resp, err
}
