package routers

import (
	"glossika/internal/controllers"
	"glossika/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type Options struct {
	DB          db.I
	RedisClient *redis.Client
}

func SetupRouter(opt Options) *gin.Engine {
	r := gin.Default()

	c := controllers.New(controllers.Options{
		DB:          opt.DB,
		RedisClient: opt.RedisClient,
	})

	v1 := r.Group("/v1")
	{
		v1.POST("/user/register", c.UserRegister)
		v1.POST("/user/login", c.UserLogin)
		v1.POST("/user/send-verification-email", c.SendVerificationEmail)
		v1.POST("/user/verify-email", c.VerifyEmail)
	}

	return r
}
