package redis

import (
	"fmt"
	"glossika/internal/config"

	"github.com/go-redis/redis"
)

func Init() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s",
			config.AppConfig.RedisHost,
			config.AppConfig.RedisPort),
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("failed to ping redis client: %w", err)
	}

	return client, nil
}
