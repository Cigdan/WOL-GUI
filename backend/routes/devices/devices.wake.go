package devices

import (
	"backend/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

func WakeDevice(c *gin.Context) {
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
}
