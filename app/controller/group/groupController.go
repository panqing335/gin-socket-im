package group

import (
	"github.com/gin-gonic/gin"
	"temp/app/entity/dto"
	"temp/app/entity/qo"
	"temp/app/service"
	util "temp/app/utils"
	"temp/app/utils/paramsBindEntity"
)

var groupService service.GroupService

// GetGroups 获取群列表
func GetGroups(c *gin.Context) {
	id := util.GetID(c)
	paginatorQo := qo.PaginatorQo{
		Page:     0,
		PageSize: 10,
	}
	paramsBindEntity.Bind(c, &paginatorQo)

	paginator := groupService.GetGroups(id, paginatorQo)
	util.Success(c, paginator)
}

// AddGroup 创建群
func AddGroup(c *gin.Context) {
	addGroupDto := dto.AddGroupDto{}
	paramsBindEntity.Bind(c, &addGroupDto)
	addGroupDto.UserId = util.GetID(c)

	groupService.AddGroup(addGroupDto)
	util.Success(c, true)
}

// JoinGroup 加入群
func JoinGroup(c *gin.Context) {
	var joinG struct {
		Id   int64  `json:"id"`
		Uuid string `json:"uuid"`
	}
	paramsBindEntity.Bind(c, &joinG)
	joinG.Id = util.GetID(c)
	groupService.JoinGroup(joinG.Id, joinG.Uuid)
	util.Success(c, true)
}
