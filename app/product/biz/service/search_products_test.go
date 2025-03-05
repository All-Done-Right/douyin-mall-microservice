package service

import (
	"github.com/All-Done-Right/douyin-mall-microservice/product/biz/dal/mysql"

	"context"
	"fmt"
	product "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
	"testing"
)

func TestSearchProducts_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewSearchProductsService(ctx)
	// init req and assert value

	req := &product.SearchProductsReq{}
	resp, err := s.Run(req)
	fmt.Println(resp)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
	// todo: edit your unit test

}
