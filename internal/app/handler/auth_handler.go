package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanedaaaa/shortl/internal/app/service"
)

func SignupHandler(c *gin.Context) {
	var userRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input data",
		})
		return
	}

	err := service.Signup(userRequest.Email, userRequest.Username, userRequest.Password)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "ok",
	})
}
