package services

import (
	"errors"
	"glossika/internal/db"
	"glossika/internal/users"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type userLoginReq struct {
	Email         string `json:"email"`
	PlainPassword string `json:"password"`
}

func (s *services) UserLogin(ctx *gin.Context) {
	var req userLoginReq
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

	verified := users.VerifyPassword(req.PlainPassword, userDB.HashedPassword)
	if !verified {
		ctx.JSON(401, gin.H{"error": "Incorrect password"})
		return
	}

	// TODO: respond JWT token
	ctx.JSON(200, gin.H{"message": "Login successful"})
}
