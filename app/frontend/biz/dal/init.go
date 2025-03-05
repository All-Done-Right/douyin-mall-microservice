package dal

import (
	"github.com/All-Done-Right/douyin-mall-microservice/app/frontend/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/app/frontend/biz/dal/redis"
)
func Init() {
	redis.Init()
	mysql.Init()
}
