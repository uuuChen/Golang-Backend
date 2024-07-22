package controllers

import (
	"glossika/internal/db"
	"glossika/internal/products"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type ControllersI interface {
	UserRegister(ctx *gin.Context)
	UserLogin(ctx *gin.Context)
	SendVerificationEmail(ctx *gin.Context)
	VerifyEmail(ctx *gin.Context)
	GetRecommendations(ctx *gin.Context)
}

type Options struct {
	DB          db.I
	RedisClient *redis.Client
}

type controllers struct {
	db          db.I
	redisClient *redis.Client
	product     products.I
}

func New(opt Options) ControllersI {
	return &controllers{
		db:          opt.DB,
		redisClient: opt.RedisClient,
		product: products.New(products.Options{
			DB:          opt.DB,
			RedisClient: opt.RedisClient,
		}),
	}
}
