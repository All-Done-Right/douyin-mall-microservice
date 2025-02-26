package core

import (
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/dal/redis"
)

func InitDB() {
	mysql.InitMysql()
	redis.InitRedis()
}
