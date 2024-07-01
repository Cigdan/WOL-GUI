package auth

import (
	"backend/utils"
	"backend/utils/auth"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user utils.UserCredentials
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "Invalid JSON format"})
		return
	}

	driver, ok := c.MustGet("driver").(*sql.DB)
	if !ok {
		c.JSON(500, gin.H{"message": "Invalid database driver"})
		return
	}

	err, code := auth.CreateUser(driver, user.Username, user.Password)
	if err != nil {
		c.JSON(code, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Successfully created new account"})
}
