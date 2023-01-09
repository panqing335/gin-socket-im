package service

import (
	"github.com/google/uuid"
	"temp/app/constants/errorCode"
	"temp/app/entity/dto"
	"temp/app/entity/qo"
	"temp/app/model"
	util "temp/app/utils"
	"temp/app/utils/paginator"
)

type GroupService struct {
}

func (g *GroupService) GetGroups(id int64, qo qo.PaginatorQo) *paginator.PaginatorCollection[map[string]any] {
	group := &model.Group{UserId: id}
	res := group.ItemsAndTotal(qo)

	paginatorCollection := paginator.NewPaginatorCollection(qo.Page, res.Items, qo.PageSize, res.Total)

	return paginatorCollection
}

func (g *GroupService) AddGroup(addGroupDto dto.AddGroupDto) {
	user := &model.User{ID: addGroupDto.UserId}
	user = util.NewResult(user.Find()).Unwrap()

	group := &model.Group{
		UUID:   uuid.New().String(),
		UserId: addGroupDto.UserId,
		Name:   addGroupDto.Name,
	}
	group = util.NewResult(group.Save()).Unwrap()

	groupMember := &model.GroupMember{
		UserId:   user.ID,
		GroupId:  group.ID,
		Nickname: user.Nickname,
		Mute:     0,
	}

	util.NewResult(groupMember.Save()).Unwrap()
}

func (g *GroupService) GetUsersByGroupUuid(groupUuid string) *[]model.UserReader {
	group := &model.Group{UUID: groupUuid}
	group = util.NewResult(group.FindByUUID()).Unwrap()

	users := util.NewResult(group.FindUsersByGroupId()).Unwrap()
	return users
}

func (g *GroupService) JoinGroup(id int64, groupUuid string) bool {
	user := &model.User{ID: id}
	user = util.NewResult(user.Find()).Unwrap()

	group := &model.Group{UUID: groupUuid}
	group = util.NewResult(group.FindByUUID()).Unwrap()

	groupMember := &model.GroupMember{UserId: user.ID, GroupId: group.ID}
	if groupMember.Exists() == true {
		util.Fail(errorCode.BAD_REQUEST, "已加入群")
	}

	groupMember.Nickname = user.Nickname
	groupMember.Mute = 0

	util.NewResult(groupMember.Save()).Unwrap()

	return true
}
