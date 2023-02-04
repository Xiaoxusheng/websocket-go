package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

/*
redis服务器
*/

var Rdb *redis.Client

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "39.98.40.255:6379",
		Password: "admin123", // no password set
		DB:       0,          // use default DB
		PoolSize: 1000,
	})
	ctx := context.Background()
	ping := client.Ping(ctx)
	fmt.Println(ping)
	Rdb = client

}
