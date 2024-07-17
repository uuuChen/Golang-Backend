package routers

import (
	"glossika/internal/db"
	"glossika/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type Options struct {
	DB          db.I
	RedisClient *redis.Client
}

func SetupRouter(opt Options) *gin.Engine {
	r := gin.Default()

	svc := services.New(services.Options{
		DB:          opt.DB,
		RedisClient: opt.RedisClient,
	})

	v1 := r.Group("/v1")
	{
		v1.POST("/user/register", svc.UserRegister)
	}

	return r
}
