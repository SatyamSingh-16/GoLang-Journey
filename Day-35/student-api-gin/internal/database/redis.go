package database

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	client := redis.NewClient(

		&redis.Options{Addr: "localhost:6379", Password: "",
			DB: 0, DialTimeout: 5 * time.Second, ReadTimeout: 3 * time.Second,
			WriteTimeout: 3 * time.Second},
	)
	ctx := context.Background()
	err := client.Ping(ctx).Err()
	if err != nil {
		log.Fatal(err)
	}
	return client
}
