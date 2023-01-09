package user

import (
	"github.com/gin-gonic/gin"
	"temp/app/entity/dto"
	"temp/app/entity/qo"
	"temp/app/model"
	"temp/app/service"
	util "temp/app/utils"
	"temp/app/utils/paramsBindEntity"
)

var userService service.UserService

// Login 登录/**
func Login(c *gin.Context) {
	loginDto := dto.LoginDto{}
	paramsBindEntity.Bind(c, &loginDto)

	token, uuid, username, nickname := userService.Login(&loginDto)

	data := struct {
		Token    string `json:"token"`
		Uuid     string `json:"uuid"`
		Username string `json:"username"`
		Nickname string `json:"nickname"`
	}{
		token,
		uuid,
		username,
		nickname,
	}

	util.Success(c, data)
}

// Register 注册用户/**
func Register(c *gin.Context) {
	registerDto := dto.RegisterDto{}
	paramsBindEntity.Bind(c, &registerDto)
	userService.Register(&registerDto)

	util.Success(c, registerDto)
}

// EditUserInfo 修改用户信息/**
func EditUserInfo(c *gin.Context) {
	var userDto dto.EditUserDto
	paramsBindEntity.Bind(c, &userDto)
	userDto.ID = util.GetID(c)

	userService.EditUserInfo(&userDto)

	util.Success(c, userDto)

}

// GetUserInfo 获取用户详情/**
func GetUserInfo(c *gin.Context) {
	//uuid := c.Query("uuid")
	uuid := util.GetUUID(c)
	util.Success(c, userService.GetUserInfo(uuid))
}

// GetUserOrGroupByName 通过名称查找群组或者用户/**
func GetUserOrGroupByName(c *gin.Context) {
	name := c.Query("name")
	user, group := userService.GetUserOrGroupByName(name)
	data := struct {
		User  *model.UserReader `json:"user"`
		Group *model.Group      `json:"group"`
	}{
		user,
		group,
	}
	util.Success(c, data)
}

// GetFriendList 获取好友列表/**
func GetFriendList(c *gin.Context) {
	id := util.GetID(c)
	paginatorQo := qo.PaginatorQo{
		Page:     0,
		PageSize: 10,
	}
	paramsBindEntity.Bind(c, &paginatorQo)
	var searchUserQo qo.SearchUserQo
	paramsBindEntity.Bind(c, &searchUserQo)
	paginator := userService.GetFriendList(id, paginatorQo, searchUserQo)

	util.Success(c, paginator)
}

// SearchUserList 查询用户/**
func SearchUserList(c *gin.Context) {
	var searchUserQo qo.SearchUserQo
	paramsBindEntity.Bind(c, &searchUserQo)
	list := userService.SearchUserList(searchUserQo)

	util.Success(c, list)
}

// AddFriend 添加好友/**
func AddFriend(c *gin.Context) {
	var addFriendDto dto.AddFriendDto
	paramsBindEntity.Bind(c, &addFriendDto)
	addFriendDto.Id = util.GetID(c)
	userService.AddFriend(&addFriendDto)

	util.Success(c, true)
}
