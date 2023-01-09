package paramsBindEntity

import (
	"github.com/gin-gonic/gin"
	"temp/app/constants/errorCode"
	util "temp/app/utils"
)

func Bind[T any](c *gin.Context, entity *T) *T {
	err := c.ShouldBind(&entity)
	if err != nil {
		util.Fail(errorCode.BAD_REQUEST, err.Error())
	}
	return entity
}
