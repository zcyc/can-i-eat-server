package redis_infrastructure

import (
	"can-i-eat/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var redisClient *redis.Client
var ctx = context.Background()

func Init() {
	// 解析配置文件
	configMap := config.Init("./config/redis/config")
	host := configMap["host"]
	port := configMap["port"]
	password := configMap["password"]
	db, _ := strconv.Atoi(configMap["database"])
	addr := fmt.Sprintf("%s:%s", host, port)
	// 连接 redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	redisClient = rdb
}

func Get() *redis.Client {
	return redisClient
}
