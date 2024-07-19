package products

import (
	"glossika/internal/db"

	"github.com/go-redis/redis"
)

type productHelper struct {
	db          db.I
	redisClient *redis.Client
}

type I interface {
	ListRecommendations() ([]db.Product, error)
}

type Options struct {
	DB          db.I
	RedisClient *redis.Client
}

func New(opt Options) I {
	return &productHelper{
		db:          opt.DB,
		redisClient: opt.RedisClient,
	}
}
