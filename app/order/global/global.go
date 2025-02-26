package global

import (
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/conf/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
)
