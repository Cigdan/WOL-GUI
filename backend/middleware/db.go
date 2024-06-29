package middleware

import (
	"backend/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func DbMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		driver, err := db.InitDB()
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"message": "couldn't connect to database"})
			c.Abort()
			return
		}
		defer driver.Close()
		c.Set("driver", driver)
		c.Next()
	}
}
