package auth

import (
	"backend/utils"
	"backend/utils/auth"
	"backend/utils/logger"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user utils.UserCredentials
	if err := c.BindJSON(&user); err != nil {
		logger.Error(user.Username + " sent invalid JSON format")
		c.JSON(400, gin.H{"message": "Invalid JSON format"})
		return
	}

	driver, ok := c.MustGet("driver").(*sql.DB)
	if !ok {
		logger.Warning("Invalid database driver")
		c.JSON(500, gin.H{"message": "Invalid database driver"})
		return
	}

	err, code := auth.CreateUser(driver, user.Username, user.Password)
	if err != nil {
		logger.Error("Couldn't create new account: " + err.Error())
		c.JSON(code, gin.H{"message": err.Error()})
		return
	}
	logger.Info("New account created: " + user.Username)
	c.JSON(200, gin.H{"message": "Successfully created new account"})
}
