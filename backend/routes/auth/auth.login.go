package auth

import (
	"backend/utils"
	"backend/utils/auth"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
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

	token, err, status := auth.Login(driver, user.Username, user.Password)
	if err != nil {
		c.JSON(status, gin.H{"message": err.Error()})
		return
	}
	c.SetCookie("token", token, 24*60*60, "/", "localhost", false, true)
	c.JSON(status, gin.H{"message": "Successfully logged in"})
}
