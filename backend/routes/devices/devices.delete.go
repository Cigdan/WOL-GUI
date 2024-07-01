package devices

import (
	"backend/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func DeleteDevice(c *gin.Context) {
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

	_, err := utils.ExecStatement(driver, "DELETE FROM device WHERE id = ? AND user_id = ?", c.Param("id"), user.ID)
	if err != nil {
		c.JSON(500, gin.H{"message": "Couldn't delete device"})
		return
	}
	c.JSON(200, gin.H{"message": "Device deleted successfully"})
}
