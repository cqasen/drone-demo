package lock

import (
	"errors"
	"github.com/ebar-go/ego/app"
	"time"
)

type RedisLock struct {
	Key string
}

func (lock *RedisLock) Lock(second int) error {
	if second == 0 {
		second = 2
	}
	res, err := app.Redis().SetNX(lock.Key, 1, time.Duration(second)*time.Second).Result()
	if err != nil || res == false {
		return errors.New("failed to lock")
	}
	return nil
}

func (lock *RedisLock) Unlock() error {
	return app.Redis().Del(lock.Key).Err()
}
