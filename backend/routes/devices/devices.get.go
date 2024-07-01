package devices

import (
	"backend/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func GetDevices(c *gin.Context) {
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

	deviceRows, err := utils.Query(driver, "SELECT * FROM device WHERE user_id = ?", user.ID)
	if err != nil {
		c.JSON(500, gin.H{"message": "Couldn't get devices", "data": nil})
		return
	}

	var devices []utils.Device
	for deviceRows.Next() {
		var device utils.Device
		err = deviceRows.Scan(&device.ID, &device.Name, &device.MacAddress, &device.IpAddress, &device.LastOnline, &device.UserID)
		if err != nil {
			c.JSON(500, gin.H{"message": "Couldn't get devices", "data": nil})
			return
		}
		devices = append(devices, device)
	}
	c.JSON(200, gin.H{"message": "Successfully fetched devices", "data": devices})
}
