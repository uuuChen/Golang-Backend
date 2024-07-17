package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

func Init() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("failed to ping redis client: %w", err)
	}

	return client, nil
}
