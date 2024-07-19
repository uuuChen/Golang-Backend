package controllers

import (
	"errors"
	"glossika/internal/db"
	"glossika/internal/users"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type userReisterReq struct {
	Email         string `json:"email"`
	PlainPassword string `json:"password"`
}

func (req *userReisterReq) toUser() *users.User {
	return &users.User{
		Email:         req.Email,
		PlainPassword: req.PlainPassword,
	}
}

func (s *controllers) UserRegister(ctx *gin.Context) {
	var req userReisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	userDB, err := s.db.FindUserByEmail(req.Email)
	if userDB != nil && userDB.Email == req.Email {
		ctx.JSON(400, gin.H{"error": "Email has been used"})
		return
	}
	if !errors.Is(err, db.NotFoundError) {
		logrus.WithError(err).Error("Failed to find user by email")
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	if !users.IsValidEmailFormat(req.Email) {
		ctx.JSON(400, gin.H{"error": "Invalid email format"})
		return
	}
	if !users.IsValidPasswordFormat(req.PlainPassword) {
		ctx.JSON(400, gin.H{"error": "Invalid password format"})
		return
	}

	userDB, err = req.toUser().ToDB()
	if err != nil {
		logrus.WithError(err).Error("Failed to transform to db format")
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err = s.db.InsertUsers([]db.User{*userDB})
	if err != nil {
		logrus.WithError(err).Error("Failed to insert users")
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err = s.sendVerificationEmail(req.Email, 60*time.Minute)
	if err != nil {
		logrus.WithError(err).Error("Failed to send verification email")
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(201, gin.H{"message": "Registration successful, please verify your email"})
}
