package devices

import (
	"backend/utils"
	"backend/utils/logger"
	"database/sql"
	"github.com/gin-gonic/gin"
	"strings"
)

func UpdateDevice(c *gin.Context) {
	user, ok := c.MustGet("userdata").(utils.UserData)
	if !ok {
		logger.Warning("Invalid userdata")
		c.JSON(500, gin.H{"message": "Invalid userdata"})
		return
	}
	var device utils.Device
	if err := c.BindJSON(&device); err != nil {
		logger.Warning(user.Username + " sent an invalid JSON format")
		c.JSON(400, gin.H{"message": "Invalid JSON format"})
		return
	}
	driver, ok := c.MustGet("driver").(*sql.DB)
	if !ok {
		logger.Warning("Invalid database driver")
		c.JSON(500, gin.H{"message": "Invalid database driver"})
		return
	}

	var replacer = strings.NewReplacer("-", ":")
	_, err := utils.ExecStatement(driver, "UPDATE device SET name = ?, mac_address = ?, ip_address = ? WHERE id = ? AND user_id = ?",
		device.Name, replacer.Replace(device.MacAddress), device.IpAddress, c.Param("id"), user.ID)
	if err != nil {
		logger.Error("Couldn't update device: " + err.Error())
		c.JSON(500, gin.H{"message": "Couldn't update device"})
		return
	}
	logger.Info("User: " + user.Username + " updated a device: " + device.Name + " with MAC: " + device.MacAddress)
	c.JSON(200, gin.H{"message": "Device updated successfully"})
}
