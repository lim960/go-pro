package common

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"time"
)

var ctx = context.Background()

var client *redis.Client

func InitRedis() {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	password := viper.GetString("redis.password")
	conn := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password, // no password set
		DB:       0,        // use default DB
	})
	client = conn
}

func Set(key string, value string) {
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

// SetWithTime seconds 秒后过期
func SetWithTime(key string, value string, seconds uint64) {
	err := client.SetEX(ctx, key, value, time.Duration(seconds)*time.Second).Err()
	if err != nil {
		panic(err)
	}
}

// SetNx 保存成功返回true  key已存在返回false
func SetNx(key, value string, seconds uint) bool {
	return client.SetNX(ctx, key, value, time.Duration(seconds)*time.Second).Val()
}

func Get(key string) string {
	value, _ := client.Get(ctx, key).Result()
	return value
}

func Del(key string) {
	client.Del(ctx, key)
}
