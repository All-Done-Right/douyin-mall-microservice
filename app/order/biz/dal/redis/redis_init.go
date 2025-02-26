package redis

import (
	"context"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/global"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func InitRedis() {
	// 初始化 Redis 连接
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     global.Config.Redis.Addr(),
			Password: global.Config.Redis.Password,
			DB:       global.Config.Redis.Database,
		},
	)
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	logrus.Infoln("successfully connecting to redis", pong)
	global.RDB = rdb

}
