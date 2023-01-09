package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"temp/app/constants/errorCode"
	_ "temp/config"
)

type BusinessException struct {
	Code    int    `json:"code"`
	Message string `json:"string"`
}

func Response(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

func Success(c *gin.Context, data interface{}) {
	Response(c, errorCode.SUCCESS, errorCode.GetMsg(errorCode.SUCCESS), data)
}

func Fail(errCode int, message string) {
	if message != "" {
		panic(BusinessException{
			Code:    errCode,
			Message: message,
		})
	} else {
		panic(BusinessException{
			Code:    errCode,
			Message: errorCode.GetMsg(errCode),
		})
	}
}
