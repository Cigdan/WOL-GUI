package main

import (
	"backend/db"
	"backend/db/auth"
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func DbMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		driver, err := db.InitDB()
		if err != nil {
			c.JSON(500, gin.H{"message": "couldn't connect to database"})
			c.Abort()
			return
		}
		defer driver.Close()
		c.Set("dbConn", driver)
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
	}))

	r.Use(DbMiddleWare())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "success"})
	})

	r.POST("/auth/login", func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"message": "Invalid JSON format"})
			return
		}

		dbConn, exists := c.Get("dbConn")
		if !exists {
			c.JSON(500, gin.H{"message": "Database connection not found"})
			return
		}

		driver, ok := dbConn.(*sql.DB)
		if !ok {
			c.JSON(500, gin.H{"message": "Invalid database driver"})
			return
		}

		token, err, status := auth.Login(driver, user.Username, user.Password)
		if err != nil {
			c.JSON(status, gin.H{"message": err.Error()})
			return
		}
		c.SetCookie("token", token, 24*60*60*1000, "/", "localhost", false, true)
		c.JSON(status, gin.H{"message": "Successfully logged in"})
	})

	r.POST("/auth/register", func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"message": "Invalid JSON format"})
			return
		}

		dbConn, exists := c.Get("dbConn")
		if !exists {
			c.JSON(500, gin.H{"message": "Database connection not found"})
			return
		}

		driver, ok := dbConn.(*sql.DB)
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

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
