package routes

import (
	"backend/routes/auth"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(authRoutes *gin.RouterGroup) {
	authRoutes.POST("/login", auth.Login)
	authRoutes.POST("/register", auth.Register)
	authRoutes.POST("/logout", auth.Logout)
	authRoutes.GET("/check", auth.CheckAuth)
}
