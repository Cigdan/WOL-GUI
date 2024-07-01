package devices

import (
	"backend/utils"
	"backend/utils/logger"
	"database/sql"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetDevices(c *gin.Context) {
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

	search := c.Query("search")
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}
	searchPattern := "%" + search + "%"
	deviceRows, err := utils.Query(driver, ""+
		"SELECT * FROM device WHERE user_id = ? AND (device.name LIKE ? OR device.ip_address LIKE ? OR device.mac_address LIKE ?) LIMIT ? OFFSET ?",
		user.ID, searchPattern, searchPattern, searchPattern, limit, offset)
	if err != nil {
		logger.Error("Couldn't get devices: " + err.Error())
		c.JSON(500, gin.H{"message": "Couldn't get devices", "data": nil})
		return
	}

	var count int
	countRow, err := utils.QueryOne(driver,
		"SELECT COUNT(*) FROM device WHERE user_id = ? AND (device.name LIKE ? OR device.ip_address LIKE ? OR device.mac_address LIKE ?)", user.ID, searchPattern, searchPattern, searchPattern)
	if err != nil {
		logger.Error("Couldn't get devices: " + err.Error())
		c.JSON(500, gin.H{"message": "Couldn't get devices", "data": nil})
		return
	}
	err = countRow.Scan(&count)
	if err != nil {
		logger.Error("Couldn't get devices: " + err.Error())
		c.JSON(500, gin.H{"message": "Couldn't get devices", "data": nil})
		return
	}

	var devices []utils.Device
	for deviceRows.Next() {
		var device utils.Device
		err = deviceRows.Scan(&device.ID, &device.Name, &device.MacAddress, &device.IpAddress, &device.LastOnline, &device.UserID)
		if err != nil {
			logger.Error("Couldn't get devices: " + err.Error())
			c.JSON(500, gin.H{"message": "Couldn't get devices", "data": nil})
			return
		}
		devices = append(devices, device)
	}
	c.JSON(200, gin.H{"message": "Successfully fetched devices", "devices": devices, "count": count})
}
