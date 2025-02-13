package main

import (
	"context"
	"net"
	"strings"

	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/dal"
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/conf"
	"github.com/All-Done-Right/douyin-mall-microservice/common/mtl"
	"github.com/All-Done-Right/douyin-mall-microservice/common/serversuite"
	"github.com/All-Done-Right/douyin-mall-microservice/common/utils"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ServiceName  = conf.GetConf().Kitex.Service
	RegisterAddr = conf.GetConf().Registry.RegistryAddress[0]
)

func main() {
	err := godotenv.Load()
	if err != nil {
		klog.Error(err.Error())
	}

	mtl.InitLog(&lumberjack.Logger{
		Filename:   conf.GetConf().Kitex.LogFileName,
		MaxSize:    conf.GetConf().Kitex.LogMaxSize,
		MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
		MaxAge:     conf.GetConf().Kitex.LogMaxAge,
	})

	mtl.InitTracing(ServiceName)
	mtl.InitMetric(ServiceName,
		conf.GetConf().Kitex.MetricsPort,
		RegisterAddr)
	p := mtl.InitTracing(ServiceName)
	defer p.Shutdown(context.Background())
	dal.Init()

	opts := kitexInit()

	svr := userservice.NewServer(new(UserServiceImpl), opts...)

	err = svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	address := conf.GetConf().Kitex.Address
	if strings.HasPrefix(address, ":") {
		localIp := utils.MustGetLocalIPv4()
		address = localIp + address
	}
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err)
	}

	opts = append(opts, server.WithServiceAddr(addr), server.WithSuite(serversuite.CommonServerSuite{CurrentServiceName: ServiceName, RegisterAddr: RegisterAddr}))
	return
}
