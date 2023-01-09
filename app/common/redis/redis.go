package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"sync"
)

var (
	myRedis   *redis.Client
	redisOnce sync.Once
)

func Setup() {
	redisOnce.Do(func() {
		myRedis = redis.NewClient(&redis.Options{
			Addr:     viper.GetString("redis.host"),
			Password: viper.GetString("redis.password"),
			DB:       viper.GetInt("redis.db"),
		})

		_, err := myRedis.Ping(context.Background()).Result()
		if err != nil {
			fmt.Println("redis链接异常：", err)
			return
		}
		fmt.Println("redis链接成功！")
	})
}
