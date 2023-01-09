package middleware

import (
	"github.com/gin-gonic/gin"
	"temp/app/constants/errorCode"
	util "temp/app/utils"
)

func SocketAuth(c *gin.Context) {
	func() {
		tokenString := c.GetHeader("Sec-Websocket-Protocol")
		if tokenString == "" {
			util.Fail(errorCode.AUTH_FAIL, "")
			c.Abort()
			return
		}
		token, claims, err := util.ParseToken(tokenString)

		if err != nil || !token.Valid {
			util.Fail(errorCode.AUTH_FAIL, "")
			c.Abort()
			return
		}
		// 从token中解析出来的数据挂载到上下文上,方便后面的控制器使用
		c.Set("Id", claims.Id)
		c.Set("UserName", claims.Username)
		c.Set("Uuid", claims.Uuid)
		c.Set("Sec-Websocket-Protocol", tokenString)
	}()
	c.Next()
}
