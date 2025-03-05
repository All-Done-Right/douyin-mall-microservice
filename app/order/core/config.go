package core

import (
	"fmt"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/global"
	"github.com/Mmx233/EnvConfig"
	"github.com/joho/godotenv"
	"github.com/pelletier/go-toml"
	"log"
	"os"
)

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	configPath := "conf/%s/config.toml"
	configPath = fmt.Sprintf(configPath, "test")
	if mode := os.Getenv("ENV"); mode != "" {
		configPath = fmt.Sprintf(configPath, mode)
	}
	tree, err := toml.LoadFile(configPath)
	if err != nil {
		log.Fatalf("Error loading TOML file: %v", err)
	}
	err = tree.Unmarshal(&global.Config)
	if err != nil {
		log.Fatalf("Error unmarshalling TOML into struct: %v", err)
	}
	EnvConfig.Load("ORDER_SERVICE_", &global.Config.ServiceInfo)
}
