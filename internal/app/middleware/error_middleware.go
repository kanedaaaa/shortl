package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanedaaaa/shortl/internal/app/errors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errs := c.Errors
		if len(errs) == 0 {
			return
		}

		err := errs[0].Err

		if customErr, ok := err.(*errors.CustomError); ok {
			if customErr.LogMsg != "" {
				log.Println(customErr.LogMsg)
			}

			c.JSON(customErr.Code, gin.H{
				"error": customErr.Message,
			})
			return
		}

		log.Fatalf("Unhandled error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
	}
}
