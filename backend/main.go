package main

import (
	"backend/db"
	"backend/db/auth"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	db.InitDB()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "success"})
	})

	r.POST("/register", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"message": "Invalid JSON format"})
			return
		}
		err, code := auth.CreateUser(user.Username, user.Password)
		if err != nil {
			c.JSON(code, gin.H{"message": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Successfully created new account"})
	})

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
