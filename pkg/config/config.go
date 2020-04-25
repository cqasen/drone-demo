package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

func InitConfig() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}
	log.Println("获取的环境变量：" + env)
	//获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.AddConfigPath(path + "/config")
	viper.SetConfigName(fmt.Sprintf("config_%s", env))
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
