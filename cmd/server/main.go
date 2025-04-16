package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kanedaaaa/shortl/internal/app/handler"
	"github.com/kanedaaaa/shortl/internal/app/middleware"
	"github.com/kanedaaaa/shortl/internal/db"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logrus.Info("Starting...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connect()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(middleware.ErrorHandler())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})
	r.GET("/health", handler.HealthHandler)

	v1 := r.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", handler.SignupHandler)
			auth.POST("/login", handler.LoginHandler)
		}

		link := v1.Group("/link")
		link.Use(middleware.AuthMiddleware())
		{
			link.POST("/shorten", handler.ShortenURLHandler)
		}
	}

	logrus.Info("Server started on :8080")
	runErr := r.Run(":8080")

	if runErr != nil {
		logrus.Fatal("failed to start: ", runErr)
	}

}
