package routes

import (
	"backend/routes/devices"
	"github.com/gin-gonic/gin"
)

func DeviceRoutes(deviceRoutes *gin.RouterGroup) {
	deviceRoutes.GET("", devices.GetDevices)
	deviceRoutes.POST("", devices.AddDevice)
	deviceRoutes.PUT("/edit/:id", devices.UpdateDevice)
	deviceRoutes.DELETE("/delete/:id", devices.DeleteDevice)
	deviceRoutes.GET("/status/:id", devices.CheckDeviceStatus)
	deviceRoutes.POST("/wake/:id", devices.WakeDevice)
}
