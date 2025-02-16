package main

import (
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/dal"
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/model"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		klog.Error(err.Error())
	}
	dal.Init()
	// mysql.DB.Create(&model.User{Email: "test@qq.com", PasswordHashed: "<password>"})

	// mysql.DB.Model(&model.User{}).Where("email = ?", "test@qq.com").Update("password_hashed", "update!!")

	// var row model.User
	// mysql.DB.Model(&model.User{}).Where("email = ?", "test@qq.com").First(&row)
	// fmt.Print(row)

	// 用Model的话会触发钩子方法，触发删除前后操作
	mysql.DB.Where("email = ?", "test@qq.com").Delete(&model.User{})
	mysql.DB.Model(&model.User{}).Where("email = ?", "test@qq.com").Delete(&model.User{})
}
