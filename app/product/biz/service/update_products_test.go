package service

import (
	"context"
	"github.com/All-Done-Right/douyin-mall-microservice/product/biz/dal/mysql"
	product "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"

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
	p.Id = 4
	p.Name = "王五"
	req := &product.UpdateProductsReq{
		Product: &p,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
