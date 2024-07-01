package main

import (
	"backend/middleware"
	"backend/routes"
	"backend/utils"
	"backend/utils/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {

	if _, err := os.Stat("./data"); os.IsNotExist(err) {
		err = os.Mkdir("./data", os.ModeDir)
		if err != nil {
			log.Fatal(err)
		}
	}

	logger.InitLogger()
	_, err := utils.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Use(middleware.EnvMiddleware())
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowCredentials = true
	r.Use(cors.New(config))
	r.Static("/assets", "../frontend/dist/assets")
	r.StaticFile("/vite.svg", "../frontend/dist/vite.svg")

	// Serve the index.html file for root and other routes
	r.GET("/", func(c *gin.Context) {
		c.File("../frontend/dist/index.html")
	})
	r.GET("/index.html", func(c *gin.Context) {
		c.File("../frontend/dist/index.html")
	})

	r.NoRoute(func(c *gin.Context) {
		c.File("../frontend/dist/index.html")
	})

	// ** Auth Routes start **
	authRoutes := r.Group("/api/auth")
	{
		authRoutes.Use(middleware.DbMiddleWare())
		routes.AuthRoutes(authRoutes)
	}
	// ** Auth Routes end **

	// ** Device Routes start **
	deviceRoutes := r.Group("/api/devices")
	{
		deviceRoutes.Use(middleware.DbMiddleWare())
		deviceRoutes.Use(middleware.AuthMiddleWare()) // Apply Auth Middleware
		routes.DeviceRoutes(deviceRoutes)
	}
	// ** Device Routes end **

	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
