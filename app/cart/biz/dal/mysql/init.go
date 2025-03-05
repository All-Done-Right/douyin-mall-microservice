package mysql

import (
	"fmt"
	"os"

	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

/*
func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		})
	DB.AutoMigrate(&model.Cart{})
	if err != nil {
		panic(err)
	}
}
*/

func Init() {
	// 1. 构建 DSN
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/cart?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
	)

	// 2. 连接数据库
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		//Logger:                 gorm.Default.LogMode(gorm.Info), // 启用日志
	})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}

	// 3. 检查数据库连接
	if err := DB.Exec("SELECT 1").Error; err != nil {
		panic("数据库连接测试失败: " + err.Error())
	}

	// 4. 自动迁移
	if err := DB.AutoMigrate(&model.Cart{}); err != nil {
		panic("自动迁移失败: " + err.Error())
	}
}

/*
func Init() {
    dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
        PrepareStmt:            true,
        SkipDefaultTransaction: true,
    })
    if err != nil {
        panic("连接数据库失败: " + err.Error())
    }

    // 自动迁移（如果使用）
    if err := DB.AutoMigrate(&model.Cart{}); err != nil {
        panic("自动迁移失败: " + err.Error())
    }

    // 或手动执行建表（如果使用）
    if err := DB.Exec(`
        CREATE TABLE IF NOT EXISTS cart (
            id         INT AUTO_INCREMENT PRIMARY KEY,
            user_id    INT NOT NULL,
            product_id INT NOT NULL,
            qty        INT NOT NULL,
            created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
        )
    `).Error; err != nil {
        panic("建表失败: " + err.Error())
    }
}
*/
