package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kanedaaaa/shortl/internal/app/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logrus.Info("Starting...")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	r.GET("/health", handler.HealthHandler)
	r.POST("/signup", handler.SignupHandler)

	err := r.Run(":8080")
	if err != nil {
		logrus.Fatal("failed to start: ", err)
	}
}
