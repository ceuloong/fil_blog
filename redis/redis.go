package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func connect() *redis.Client {
	// 连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	return rdb
}

func SetRedis(key string, value interface{}, expiration time.Duration) error {
	rdb := connect()
	defer rdb.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 设置值
	err := rdb.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetRedis(key string) (string, error) {
	rdb := connect()
	defer rdb.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 获取值
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
