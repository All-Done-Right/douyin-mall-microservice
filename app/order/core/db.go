package core


import (
	redis "github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/dal/cache"
	mysql "github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/dal/repo"
)



 func InitDB() {
 	mysql.InitMysql()
 	redis.InitRedis()
 }
