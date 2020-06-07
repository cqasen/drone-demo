package lock

import "github.com/spf13/viper"

var Prefix = viper.GetString("redis.prefix")

type ILock interface {
	Lock(second int) error
	Unlock() error
}

func GetRedisLock(key string) ILock {
	return &RedisLock{Key: Prefix + key}
}
