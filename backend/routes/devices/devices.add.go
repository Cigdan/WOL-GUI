package devices

import (
	"backend/utils"
	"backend/utils/logger"
	"database/sql"
	"github.com/gin-gonic/gin"
	"strings"
)

func AddDevice(c *gin.Context) {
	var device utils.Device
	user, ok := c.MustGet("userdata").(utils.UserData)
	if !ok {
		logger.Warning("Invalid userdata")
		c.JSON(500, gin.H{"message": "Invalid userdata"})
		return
	}

	if err := c.BindJSON(&device); err != nil {
		logger.Warning(user.Username + " sent a Invalid JSON format")
		c.JSON(400, gin.H{"message": "Invalid JSON format"})
		return
	}

	if device.IpAddress != nil && *device.IpAddress == "" {
		device.IpAddress = nil
	}

	driver, ok := c.MustGet("driver").(*sql.DB)
	if !ok {
		logger.Warning("Invalid database driver")
		c.JSON(500, gin.H{"message": "Invalid database driver"})
		return
	}

	var replacer = strings.NewReplacer("-", ":")
	_, err := utils.ExecStatement(driver,
		"INSERT INTO device (name, mac_address, ip_address, user_id) VALUES (?, ?, ?, ?)",
		device.Name, replacer.Replace(device.MacAddress), device.IpAddress, user.ID)
	if err != nil {
		logger.Error("Couldn't add device: " + err.Error())
		c.JSON(500, gin.H{"message": "Couldn't add device"})
		return
	}

	logger.Info("User: " + user.Username + " added a device: " + device.Name + " with MAC: " + device.MacAddress)
	c.JSON(201, gin.H{"message": "Device added successfully"})
}
