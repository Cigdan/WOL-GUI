package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func EnvMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		c.Next()
	}
}
