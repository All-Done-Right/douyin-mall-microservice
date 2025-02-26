package dal

import (
	"app/product/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
