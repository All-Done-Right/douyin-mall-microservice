package main

import (
	"net"
	"time"

	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/dal"
	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/conf"
	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/rpc"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/All-Done-Right/douyin-mall-microservice/common/serversuite"
	"github.com/All-Done-Right/douyin-mall-microservice/common/mtl"
)

var (
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
)

func main() {
	_ = godotenv.Load()
	mtl.InitMetric(ServiceName, conf.GetConf().Kitex.MetricsPort, RegistryAddr)
	dal.Init()
	rpc.InitClient()

	opts := kitexInit()
	svr := cartservice.NewServer(new(CartServiceImpl), opts...)
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
	opts = append(opts, server.WithServiceAddr(addr), server.WithSuite(serversuite.CommonServerSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	/*
		r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
		if err != nil {
			klog.Fatal(err)
		}
	*/
	/*
		// service info
		opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: conf.GetConf().Kitex.Service,
		}), server.WithRegistry(r))
	*/
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
