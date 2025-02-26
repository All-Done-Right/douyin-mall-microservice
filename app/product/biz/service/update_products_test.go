package service

import (
	"app/product/biz/dal/mysql"
	"context"
	"github.com/joho/godotenv"
	product "rpc_gen/kitex_gen/product"
	"testing"
)

func TestUpdateProducts_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	// init req and assert value

	s := NewUpdateProductsService(ctx)
	// init req and assert value
	var p product.Product
	p.Id = 3
	p.Name = "张三"
	req := &product.UpdateProductsReq{
		Product: &p,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
