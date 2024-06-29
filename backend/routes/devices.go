package routes

import (
	"backend/db"
	"backend/middleware"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Device struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	MacAddress string     `json:"mac_address"`
	LastOnline *time.Time `json:"last_online"`
	UserID     int        `json:"user_id"`
}

func DeviceRoutes(deviceRoutes *gin.RouterGroup) {
	// Get Devices Route
	deviceRoutes.GET("", func(c *gin.Context) {
		driver, ok := c.MustGet("driver").(*sql.DB)
		if !ok {
			c.JSON(500, gin.H{"message": "Invalid database driver"})
			return
		}

		user, ok := c.MustGet("userdata").(middleware.UserData) // Use middleware.UserData
		if !ok {
			c.JSON(500, gin.H{"message": "Invalid userdata"})
			return
		}

		deviceRows, err := db.Query(driver, "SELECT * FROM device WHERE user_id = ?", user.ID)
		if err != nil {
			c.JSON(500, gin.H{"message": "Couldn't get devices", "data": nil})
			return
		}

		var devices []Device
		for deviceRows.Next() {
			var device Device
			err = deviceRows.Scan(&device.ID, &device.Name, &device.MacAddress, &device.LastOnline, &device.UserID)
			if err != nil {
				fmt.Println(err)
				c.JSON(500, gin.H{"message": "Couldn't get devices", "data": nil})
				return
			}
			devices = append(devices, device)
		}
		c.JSON(200, gin.H{"message": "Successfully fetched devices", "data": devices})
	})
}
