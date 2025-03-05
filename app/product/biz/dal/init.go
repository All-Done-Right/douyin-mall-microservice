package dal

import (
	"github.com/All-Done-Right/douyin-mall-microservice/product/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
