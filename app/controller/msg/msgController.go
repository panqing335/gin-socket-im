package msg

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"os"
	"strings"
	"temp/app/entity/dto"
	"temp/app/service"
	util "temp/app/utils"
	"temp/app/utils/paramsBindEntity"
)

// GetMsg 获取消息/**
func GetMsg(c *gin.Context) {
	var reqDto dto.MsgReqDto
	paramsBindEntity.Bind(c, &reqDto)
	reqDto.ID = util.GetID(c)

	util.Logger().Info("messageRequest params: ", reqDto)

	messageArr := service.NewMessageService().GetMessage(reqDto)

	util.Success(c, messageArr)
}

// GetFile 前端通过文件名称获取文件流，显示文件/**
func GetFile(c *gin.Context) {
	fileName := c.Param("fileName")
	file, _ := os.ReadFile(fileName)
	c.Writer.Write(file)
}

// SaveFile 上传文件/**
func SaveFile(c *gin.Context) {
	namePrefix := uuid.New().String()

	file, _ := c.FormFile("file")

	fileName := file.Filename
	index := strings.LastIndex(fileName, ".")
	suffix := fileName[index:]

	newFileName := namePrefix + suffix

	c.SaveUploadedFile(file, viper.GetString("upload.staticPath")+newFileName)

	userDto := dto.EditUserDto{
		ID:     util.GetID(c),
		Avatar: newFileName,
	}

	userService := service.UserService{}
	userService.EditUserInfo(&userDto)

	util.Success(c, newFileName)
}
