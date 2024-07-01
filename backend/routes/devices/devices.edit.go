package devices

import (
	"backend/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"strings"
)

func UpdateDevice(c *gin.Context) {
	var device utils.Device
	if err := c.BindJSON(&device); err != nil {
		c.JSON(400, gin.H{"message": "Invalid JSON format"})
		return
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
	_, err := utils.ExecStatement(driver, "UPDATE device SET name = ?, mac_address = ?, ip_address = ? WHERE id = ? AND user_id = ?",
		device.Name, replacer.Replace(device.MacAddress), device.IpAddress, c.Param("id"), user.ID)
	if err != nil {
		c.JSON(500, gin.H{"message": "Couldn't update device"})
		return
	}
	c.JSON(200, gin.H{"message": "Device updated successfully"})
}
