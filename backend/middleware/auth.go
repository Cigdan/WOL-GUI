package middleware

import (
	"backend/utils"
	"backend/utils/auth"
	"backend/utils/logger"
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			logger.Warning("No Token has been provided")
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		claims, err, status := auth.ValidateToken(token)
		if err != nil {
			c.JSON(status, gin.H{"message": err.Error()})
			c.Abort()
			return
		}
		userData := utils.UserData{Username: claims.Username, ID: claims.ID}
		c.Set("userdata", userData)
		c.Next()
	}
}
