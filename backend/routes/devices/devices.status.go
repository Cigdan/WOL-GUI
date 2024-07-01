package devices

import (
	"backend/utils"
	"backend/utils/logger"
	"database/sql"
	"github.com/gin-gonic/gin"
	probing "github.com/prometheus-community/pro-bing"
	"time"
)

func CheckDeviceStatus(c *gin.Context) {
	driver, ok := c.MustGet("driver").(*sql.DB)
	if !ok {
		logger.Warning("Invalid database driver")
		c.JSON(500, gin.H{"message": "Invalid database driver", "status": -1})
		return
	}

	user, ok := c.MustGet("userdata").(utils.UserData)
	if !ok {
		logger.Warning("Invalid userdata")
		c.JSON(500, gin.H{"message": "Invalid userdata", "status": -1})
		return
	}

	var ipAddress *string
	deviceRow, err := utils.QueryOne(driver, "SELECT ip_address FROM device WHERE id = ? AND user_id = ?", c.Param("id"), user.ID)
	if err != nil {
		logger.Warning("Couldn't get device status: " + err.Error())
		c.JSON(500, gin.H{"message": "Couldn't get device status", "status": -1})
		return
	}
	err = deviceRow.Scan(&ipAddress)
	if err != nil {
		logger.Warning("Couldn't get device status: " + err.Error())
		c.JSON(500, gin.H{"message": "Couldn't get device status", "status": -1})
		return
	}
	if ipAddress == nil {
		c.JSON(200, gin.H{"message": "Device is offline", "status": 0})
		return
	}

	pinger, err := probing.NewPinger(*ipAddress)
	if err != nil {
		logger.Warning("Couldn't get device status: " + err.Error())
		c.JSON(500, gin.H{"message": "Couldn't get device status", "status": -1})
		return
	}
	pinger.SetPrivileged(true)
	pinger.Timeout = 5 * time.Second
	pinger.Count = 1
	err = pinger.Run()
	if err != nil {
		logger.Warning("Couldn't get device status: " + err.Error())
		c.JSON(500, gin.H{"message": "Couldn't get device status", "status": -1})
		return
	}
	if pinger.Statistics().PacketsRecv == 0 {
		c.JSON(200, gin.H{"message": "Device is offline", "status": 0})
		return
	}
	c.JSON(200, gin.H{"message": "Device is online", "status": 1})
}
