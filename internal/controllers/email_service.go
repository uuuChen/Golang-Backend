package controllers

import (
	"errors"
	"fmt"
	"glossika/internal/db"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type sendVerificationEmailReq struct {
	Email string `json:"email"`
}

func (s *controllers) SendVerificationEmail(ctx *gin.Context) {
	var req sendVerificationEmailReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	userDB, err := s.db.FindUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, db.NotFoundError) {
			ctx.JSON(400, gin.H{"error": "Email has not been registered"})
			return
		} else {
			logrus.WithError(err).Error("Failed to find user by email")
			ctx.JSON(500, gin.H{"error": "Internal server error"})
			return
		}
	}

	if userDB.IsEmailVerified {
		ctx.JSON(409, gin.H{"error": "Email has already been verified"})
		return
	}

	overLimit, err := s.isOverRateLimit(req.Email)
	if err != nil {
		logrus.WithError(err).Error("Failed to check rate limit")
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	if overLimit {
		ctx.JSON(429, gin.H{"error": "Too many requests. Please try again later"})
		return
	}

	code := generateVerificationCode()

	err = s.sendVerificationEmail(req.Email, code, 60*time.Minute)
	if err != nil {
		logrus.WithError(err).Error("Failed to send verification email")
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err = s.setRateLimit(req.Email, 1*time.Minute)
	if err != nil {
		logrus.WithError(err).Error("Failed to set rate limit")
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(200, gin.H{
		"message":           "Verification email is sent",
		"verification_code": code,
	})
}

type verifyEmailReq struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func (s *controllers) VerifyEmail(ctx *gin.Context) {
	var req verifyEmailReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	userDB, err := s.db.FindUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, db.NotFoundError) {
			ctx.JSON(400, gin.H{"error": "Email has not been registered"})
			return
		} else {
			logrus.WithError(err).Error("Failed to find user by email")
			ctx.JSON(500, gin.H{"error": "Internal server error"})
			return
		}
	}

	if userDB.IsEmailVerified {
		ctx.JSON(200, gin.H{"message": "Email verified"})
		return
	}

	storedCode, err := s.redisClient.Get(req.Email).Result()
	if err == redis.Nil {
		ctx.JSON(401, gin.H{"error": "Invalid or expired code"})
		return
	} else if err != nil {
		logrus.WithError(err).Error("Failed to get code from Redis")
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	if storedCode != req.Code {
		ctx.JSON(401, gin.H{"error": "Incorrect verification code"})
		return
	}

	s.db.UpdateUser(userDB.ID, map[string]interface{}{
		db.UserColumnIsEmailVerified.ToString(): true,
	})

	ctx.JSON(200, gin.H{"message": "Email verified"})
}

func (s *controllers) isOverRateLimit(key string) (bool, error) {
	key = fmt.Sprintf("rate_limit:%s", key)
	ttl, err := s.redisClient.TTL(key).Result()
	if err != nil && err != redis.Nil {
		return false, fmt.Errorf("failed to get TTL from redis: %w", err)
	}
	return ttl > 0, nil
}

func (s *controllers) setRateLimit(key string, limitTime time.Duration) error {
	key = fmt.Sprintf("rate_limit:%s", key)
	err := s.redisClient.Set(key, "1", limitTime).Err()
	if err != nil {
		return fmt.Errorf("failed to set rate limit key in redis: %w", err)
	}
	return nil
}

func (s *controllers) sendVerificationEmail(email string, code string, expiredTime time.Duration) error {
	logrus.Infof("Send email to \"%s\" with code \"%s\"\n", email, code)

	err := s.redisClient.Set(email, code, expiredTime).Err()
	if err != nil {
		return fmt.Errorf("failed to set verification code in redis: %w", err)
	}

	return nil
}

func generateVerificationCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06d", r.Intn(1000000))
}
