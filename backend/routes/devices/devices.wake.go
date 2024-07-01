package devices

import (
	"backend/utils"
	"backend/utils/logger"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

func WakeDevice(c *gin.Context) {
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

	var macAddress string
	deviceRow, err := utils.QueryOne(driver, "SELECT mac_address FROM device WHERE id = ? AND user_id = ?", c.Param("id"), user.ID)
	if err != nil {
		logger.Error("Query Error: " + err.Error())
		fmt.Println(err)
		c.JSON(500, gin.H{"message": "Couldn't wake device"})
		return
	}
	err = deviceRow.Scan(&macAddress)
	if err != nil {
		logger.Error("Scan Error: " + err.Error())
		c.JSON(500, gin.H{"message": "Couldn't wake device"})
		return
	}

	packet, err, status := utils.GeneratePacket(macAddress)
	if err != nil {
		logger.Error("Packet Error: " + err.Error())
		c.JSON(status, gin.H{"message": err.Error()})
		return
	}
	err = utils.SendPacket(packet)
	if err != nil {
		logger.Error("Send Error: " + err.Error())
		c.JSON(500, gin.H{"message": "Couldn't sent packet"})
		return
	}
	logger.Info("Packet has been sent to: " + macAddress)
	c.JSON(200, gin.H{"message": "Packet has been sent"})
}
