package main

import (
	"fmt"
	"github.com/All-Done-Right/douyin-mall-microservice/product/biz/dal"
	"github.com/All-Done-Right/douyin-mall-microservice/product/conf"
	s "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/registry-nacos/v2/registry"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net"
	"time"
)

func main() {
	err2 := godotenv.Load("./app/product/.env")
	if err2 != nil {
		log.Fatal(err2)
		log.Fatal("Error loading .env file")
	}
	dal.Init()

	opts := kitexInit()

	svr := s.NewServer(new(ProductCatalogServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	fmt.Println("Nacos服务注册")
	//Nacos服务注册
	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		panic(err)
	}

	//其他软件的服务注册.......
	//r, err := consul.NewConsulRegister("127.0.0.1:8500")
	//if err != nil {
	//	log.Fatal(err)
	//}

	opts = append(opts, server.WithRegistry(r))
	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
