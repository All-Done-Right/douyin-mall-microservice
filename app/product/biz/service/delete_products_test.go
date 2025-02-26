package service

import (
	"app/product/biz/dal/mysql"
	"context"
	"github.com/joho/godotenv"
	product "rpc_gen/kitex_gen/product"
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
