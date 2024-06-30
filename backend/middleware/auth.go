package middleware

import (
	"backend/utils/auth"
	"github.com/gin-gonic/gin"
)

type UserData struct {
	Username string `json:"username"`
	ID       int    `json:"id"`
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
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
		userData := UserData{Username: claims.Username, ID: claims.ID}
		c.Set("userdata", userData)
		c.Next()
	}
}
