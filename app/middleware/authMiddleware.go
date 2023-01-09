package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkeridea/go-extend/exunicode/exutf8"
	"temp/app/constants/errorCode"
	util "temp/app/utils"
)

func Auth(c *gin.Context) {
	func() {
		tokenString := c.GetHeader("Authorization")
		tokenString = exutf8.RuneSubString(tokenString, 7, 0)
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
	}()
	c.Next()
}
