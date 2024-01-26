package initializers

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		fmt.Println("cannot connect redis db")
	}
	fmt.Println("Redis connected ")

	return rdb
}
