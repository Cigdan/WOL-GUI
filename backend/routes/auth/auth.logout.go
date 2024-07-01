package auth

import (
	"backend/utils/logger"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	logger.Info("User logged out")
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"message": "Successfully logged out"})
}
