package config

import "os"

func GetEnv() string {
	//加载配置
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}
	return env
}
