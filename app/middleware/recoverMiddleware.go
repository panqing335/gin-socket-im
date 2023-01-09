package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	util "temp/app/utils"
)

func Recover(c *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("panic: %v\n", err)
			util.Logger().Error("error:", err)

			switch v := err.(type) {
			case error:
				c.JSON(400, gin.H{
					"code":    400,
					"message": err.(string),
				})
			case util.BusinessException:
				c.JSON(400, gin.H{
					"code":    v.Code,
					"message": v.Message,
				})
			default:
				c.JSON(500, gin.H{
					"code":    500,
					"message": err.(string),
				})
			}
			c.Abort()
		}
	}()

	c.Next()
}
