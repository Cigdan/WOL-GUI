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
		c.Set("driver", driver)
		c.Next()
	}
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
		c.Set("claims", claims)
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

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "success"})
	})

	// ** Auth Routes start **
	authRoutes := r.Group("/auth")
	{
		authRoutes.Use(DbMiddleWare())

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
			c.SetCookie("token", token, 24*60*60*1000, "/", "localhost", false, true)
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
	}
	// ** Auth Routes end **

	// ** Device Routes start **
	deviceRoutes := r.Group("/devices")
	{
		deviceRoutes.Use(DbMiddleWare())
		deviceRoutes.Use(AuthMiddleWare())

		// Get Devices Route
		deviceRoutes.GET("/", func(c *gin.Context) {
			driver, ok := c.MustGet("driver").(*sql.DB)
			if !ok {
				c.JSON(500, gin.H{"message": "Invalid database driver"})
				return
			}
			claims, ok := c.MustGet("claims").(*Claims)
			if !ok {
				c.JSON(500, gin.H{"message": "Invalid claims"})
				return
			}
			devices, err := db.Query(driver, "SELECT * FROM device WHERE user_id = ?", claims.Id)
			if err != nil {
				c.JSON(500, gin.H{"message": "Couldn't get devices", "data": nil})
				return
			}
			c.JSON(200, gin.H{"message": "Successfully fetched devices", "data": devices})
		})
	}
	// ** Device Routes end **

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
