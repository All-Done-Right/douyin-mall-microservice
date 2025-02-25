package dal

import (
	"GoMall/app/product/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
