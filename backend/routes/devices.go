package routes

import (
	"backend/middleware"
	"backend/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus-community/pro-bing"
	"strings"
	"time"
)

type Device struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	MacAddress string     `json:"mac_address"`
	IpAddress  *string    `json:"ip_address"`
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

		user, ok := c.MustGet("userdata").(middleware.UserData)
		if !ok {
			c.JSON(500, gin.H{"message": "Invalid userdata"})
			return
		}

		deviceRows, err := utils.Query(driver, "SELECT * FROM device WHERE user_id = ?", user.ID)
		if err != nil {
			c.JSON(500, gin.H{"message": "Couldn't get devices", "data": nil})
			return
		}

		var devices []Device
		for deviceRows.Next() {
			var device Device
			err = deviceRows.Scan(&device.ID, &device.Name, &device.MacAddress, &device.IpAddress, &device.LastOnline, &device.UserID)
			if err != nil {
				c.JSON(500, gin.H{"message": "Couldn't get devices", "data": nil})
				return
			}
			devices = append(devices, device)
		}
		c.JSON(200, gin.H{"message": "Successfully fetched devices", "data": devices})
	})

	// Add Device Route
	deviceRoutes.POST("", func(c *gin.Context) {
		var device Device
		if err := c.BindJSON(&device); err != nil {
			c.JSON(400, gin.H{"message": "Invalid JSON format"})
			return
		}

		// Set IpAddress to nil if it's an empty string
		if device.IpAddress != nil && *device.IpAddress == "" {
			device.IpAddress = nil
		}

		driver, ok := c.MustGet("driver").(*sql.DB)
		if !ok {
			c.JSON(500, gin.H{"message": "Invalid database driver"})
			return
		}

		user, ok := c.MustGet("userdata").(middleware.UserData)
		if !ok {
			c.JSON(500, gin.H{"message": "Invalid userdata"})
			return
		}

		var replacer = strings.NewReplacer("-", ":")
		_, err := utils.ExecStatement(driver,
			"INSERT INTO device (name, mac_address, ip_address, user_id) VALUES (?, ?, ?, ?)",
			device.Name, replacer.Replace(device.MacAddress), device.IpAddress, user.ID)
		if err != nil {
			c.JSON(500, gin.H{"message": "Couldn't add device"})
			return
		}

		c.JSON(201, gin.H{"message": "Device added successfully"})
	})

	// Update Device Route
	deviceRoutes.PUT("/edit/:id", func(c *gin.Context) {
		var device Device
		if err := c.BindJSON(&device); err != nil {
			c.JSON(400, gin.H{"message": "Invalid JSON format"})
			return
		}
		driver, ok := c.MustGet("driver").(*sql.DB)
		if !ok {
			c.JSON(500, gin.H{"message": "Invalid database driver"})
			return
		}

		user, ok := c.MustGet("userdata").(middleware.UserData)
		if !ok {
			c.JSON(500, gin.H{"message": "Invalid userdata"})
			return
		}

		var replacer = strings.NewReplacer("-", ":")
		_, err := utils.ExecStatement(driver, "UPDATE device SET name = ?, mac_address = ?, ip_address = ? "+
			"WHERE id = ? AND user_id = ?",
			device.Name, replacer.Replace(device.MacAddress), device.IpAddress, c.Param("id"), user.ID)
		if err != nil {
			c.JSON(500, gin.H{"message": "Couldn't update device"})
			return
		}
		c.JSON(200, gin.H{"message": "Device updated successfully"})
	})

	// Delete Device Route
	deviceRoutes.DELETE("/delete/:id", func(c *gin.Context) {
		driver, ok := c.MustGet("driver").(*sql.DB)
		if !ok {
			c.JSON(500, gin.H{"message": "Invalid database driver"})
			return
		}

		user, ok := c.MustGet("userdata").(middleware.UserData)
		if !ok {
			c.JSON(500, gin.H{"message": "Invalid userdata"})
			return
		}

		_, err := utils.ExecStatement(driver, "DELETE FROM device WHERE id = ? AND user_id = ?", c.Param("id"), user.ID)
		if err != nil {
			c.JSON(500, gin.H{"message": "Couldn't delete device"})
			return
		}
		c.JSON(200, gin.H{"message": "Device deleted successfully"})
	})

	// Check Device Status
	deviceRoutes.GET("/status/:id", func(c *gin.Context) {
		driver, ok := c.MustGet("driver").(*sql.DB)
		if !ok {
			c.JSON(500, gin.H{"message": "Invalid database driver", "status": -1})
			return
		}

		user, ok := c.MustGet("userdata").(middleware.UserData)
		if !ok {
			c.JSON(500, gin.H{"message": "Invalid userdata", "status": -1})
			return
		}

		var ipAddress *string
		deviceRow, err := utils.QueryOne(driver, "SELECT ip_address FROM device WHERE id = ? AND user_id = ?", c.Param("id"), user.ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"message": "Couldn't get device status", "status": -1})
			return
		}
		err = deviceRow.Scan(&ipAddress)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"message": "Couldn't get device status", "status": -1})
			return
		}
		if ipAddress == nil {
			c.JSON(200, gin.H{"message": "Device is offline", "status": 0})
			return
		}

		pinger, err := probing.NewPinger(*ipAddress)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"message": "Couldn't get device status", "status": -1})
			return
		}
		pinger.SetPrivileged(true)
		pinger.Count = 1
		err = pinger.Run()
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"message": "Couldn't get device status", "status": -1})
			return
		}
		c.JSON(200, gin.H{"message": "Device is online", "status": 1})
	})

	// Wake Device Route
	deviceRoutes.POST("/wake/:id", func(c *gin.Context) {
		driver, ok := c.MustGet("driver").(*sql.DB)
		if !ok {
			c.JSON(500, gin.H{"message": "Invalid database driver"})
			return
		}

		user, ok := c.MustGet("userdata").(middleware.UserData)
		if !ok {
			c.JSON(500, gin.H{"message": "Invalid userdata"})
			return
		}

		var macAddress string
		deviceRow, err := utils.QueryOne(driver, "SELECT mac_address FROM device WHERE id = ? AND user_id = ?", c.Param("id"), user.ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"message": "Couldn't wake device"})
			return
		}
		err = deviceRow.Scan(&macAddress)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"message": "Couldn't wake device"})
			return
		}

		packet, err, status := utils.GeneratePacket(macAddress)
		if err != nil {
			fmt.Println(err)
			c.JSON(status, gin.H{"message": err.Error()})
			return
		}
		err = utils.SendPacket(packet)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"message": "Couldn't sent packet"})
			return
		}
		c.JSON(200, gin.H{"message": "Packet has been sent"})
	})
}
