package rpc

import (
	"sync"

	//"github.com/cloudwego/biz-demo/gomall/app/cart/conf"
	//cartutils "github.com/cloudwego/biz-demo/gomall/app/cart/utils"
	//"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"

	//"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/service"
	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/conf"
	cartutils "github.com/All-Done-Right/douyin-mall-microservice/app/cart/utils"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/biz-demo/gomall/common/clientsuite"
	"github.com/cloudwego/kitex/client"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	ServiceName   string
	RegistryAddr  string
	err           error
)

func InitClient() {
	once.Do(func() {
		RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
		ServiceName = conf.GetConf().Kitex.Service
		initProductClient()
	})
}

func initProductClient() {
	// 初始化ProductClient
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartutils.MustHandleError(err)
}
