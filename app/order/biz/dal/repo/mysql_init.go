package repo

import (
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitMysql() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: global.Config.Mysql.DSN(),
	}), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	global.DB = db

}
