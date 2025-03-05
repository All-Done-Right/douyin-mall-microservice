package rpc

import (
	"sync"

	//"github.com/All-Done-Right/douyin-mall-microservice/app/cart/conf"
	//cartutils "github.com/All-Done-Right/douyin-mall-microservice/app/cart/utils"
	//"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product/productcatalogservice"

	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/conf"
	cartutils "github.com/All-Done-Right/douyin-mall-microservice/app/cart/utils"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
)

func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartutils.MustHandleError(err)
}
