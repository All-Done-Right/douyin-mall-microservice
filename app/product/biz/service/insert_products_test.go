package service

import (
	"context"
	product "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product"
	"testing"
)

func TestInsertProducts_Run(t *testing.T) {
	ctx := context.Background()
	s := NewInsertProductsService(ctx)
	// init req and assert value

	req := &product.InsertProductsReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
