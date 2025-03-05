package dal

import (
	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/dal/mysql"
	//"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/dal/redis"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
