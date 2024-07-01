package auth

import (
	"backend/utils/auth"
	"github.com/gin-gonic/gin"
)

func CheckAuth(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, gin.H{"message": "Unauthorized"})
		return
	}
	_, err, _ = auth.ValidateToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Authorized"})
}
