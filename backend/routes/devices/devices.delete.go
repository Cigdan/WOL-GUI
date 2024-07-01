package devices

import (
	"backend/utils"
	"backend/utils/logger"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func DeleteDevice(c *gin.Context) {
	driver, ok := c.MustGet("driver").(*sql.DB)
	if !ok {
		logger.Warning("Invalid database driver")
		c.JSON(500, gin.H{"message": "Invalid database driver"})
		return
	}

	user, ok := c.MustGet("userdata").(utils.UserData)
	if !ok {
		logger.Warning("Invalid userdata")
		c.JSON(500, gin.H{"message": "Invalid userdata"})
		return
	}

	_, err := utils.ExecStatement(driver, "DELETE FROM device WHERE id = ? AND user_id = ?", c.Param("id"), user.ID)
	if err != nil {
		logger.Error(user.Username + " couldn't delete device: " + err.Error())
		c.JSON(500, gin.H{"message": "Couldn't delete device"})
		return
	}
	logger.Info("User: " + user.Username + " deleted a device: " + c.Param("id"))
	c.JSON(200, gin.H{"message": "Device deleted successfully"})
}
