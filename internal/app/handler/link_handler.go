package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanedaaaa/shortl/internal/app/service"
)

func ShortenURLHandler(c *gin.Context) {
	userIDValue, exists := c.Get("userID")

	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userID := userIDValue.(uint)

	var userRequest struct {
		URL string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input data",
		})
		return
	}

	shortened, err := service.ShortenURL(userID, userRequest.URL)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "ok",
		"data":    shortened,
	})
}

func GetURLHandler(c *gin.Context) {
	userIDValue, exists := c.Get("userID")

	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userID := userIDValue.(uint)

	urls, err := service.GetURL(userID)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    urls,
	})

}
