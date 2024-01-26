package initializers

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

var Ctx = context.Background()

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("$REDIS_URI"),
		Password: os.Getenv("$REDIS_PASS"),
		DB:       0,
	})
	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		fmt.Println("cannot connect redis client", err)
	} else {
		fmt.Println("REDIS client: OK")
	}

	return rdb
}
