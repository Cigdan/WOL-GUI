package routes

import (
	"backend/utils/auth"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterRoutes(authRoutes *gin.RouterGroup) {
	// Login Route
	authRoutes.POST("/login", func(c *gin.Context) {
		var user User
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
	})

	// Register Route
	authRoutes.POST("/register", func(c *gin.Context) {
		var user User
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
	})

	// Logout Route
	authRoutes.POST("/logout", func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "localhost", false, true)
		c.JSON(200, gin.H{"message": "Successfully logged out"})
	})

	// Check if user is logged in
	authRoutes.GET("/check", func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			return
		}
		_, err, _ = auth.ValidateToken(token)
		if err != nil {
			c.JSON(401, gin.H{"message": err.Error()})
		}
		c.JSON(200, gin.H{"message": "Authorized"})
	})
}
