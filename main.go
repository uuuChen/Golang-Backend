package main

import (
	"glossika/internal/db"
	"glossika/internal/redis"
	"glossika/internal/routers"

	"github.com/sirupsen/logrus"
)

func main() {
	db, err := db.Init()
	if err != nil {
		logrus.WithError(err).Panic("Failed to init db")
	}

	redisClient, err := redis.Init()
	if err != nil {
		logrus.WithError(err).Panic("Failed to init redis client")
	}

	r := routers.SetupRouter(routers.Options{
		DB:          db,
		RedisClient: redisClient,
	})
	r.Run(":8080")
}
