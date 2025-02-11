package dal

import (
	"github.com/All-Done-Right/douyin-mall-microservice/app/auth/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
