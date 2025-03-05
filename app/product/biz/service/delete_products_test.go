package service

import (
	"context"
	"github.com/All-Done-Right/douyin-mall-microservice/product/biz/dal/mysql"
	product "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
	"testing"
)

func TestDeleteProducts_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewDeleteProductsService(ctx)
	// init req and assert value

	req := &product.DeleteProductsReq{
		Id: 3,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
