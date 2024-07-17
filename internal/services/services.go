package services

import (
	"glossika/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type ServicesI interface {
	UserRegister(ctx *gin.Context)
}

type Options struct {
	DB          db.I
	RedisClient *redis.Client
}

type services struct {
	db          db.I
	redisClient *redis.Client
}

func New(opt Options) ServicesI {
	return &services{
		db:          opt.DB,
		redisClient: opt.RedisClient,
	}
}
