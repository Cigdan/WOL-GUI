package devices

import (
	"backend/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"strings"
)

func AddDevice(c *gin.Context) {
	var device utils.Device
	if err := c.BindJSON(&device); err != nil {
		c.JSON(400, gin.H{"message": "Invalid JSON format"})
		return
	}

	if device.IpAddress != nil && *device.IpAddress == "" {
		device.IpAddress = nil
	}

	driver, ok := c.MustGet("driver").(*sql.DB)
	if !ok {
		c.JSON(500, gin.H{"message": "Invalid database driver"})
		return
	}

	user, ok := c.MustGet("userdata").(utils.UserData)
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
}
