package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const redisKeyRecommendations = "recommendations"

func (s *controllers) GetRecommendations(ctx *gin.Context) {
	recommendations, err := s.product.ListRecommendations()
	if err != nil {
		logrus.WithError(err).Error("Failed to list recommendations")
		ctx.JSON(500, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"recommendations": recommendations})
}
