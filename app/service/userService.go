package service

import (
	"github.com/google/uuid"
	"strconv"
	"temp/app/constants/errorCode"
	"temp/app/entity/dto"
	"temp/app/entity/qo"
	"temp/app/model"
	util "temp/app/utils"
	"temp/app/utils/paginator"
)

type UserService struct {
}

func (u *UserService) Login(loginDto *dto.LoginDto) (token, uuid, username, nickname string) {
	util.Logger().Debug("loginDto", loginDto)
	user := &model.User{Username: loginDto.Username}
	user = util.NewResult(user.FindByUsername()).UnwrapOr(errorCode.LOGIN_ERROR, "")
	util.Logger().Debug("user", user)
	util.NewResult(util.CheckPassword(loginDto.Password, user.PasswordSalt, user.Password)).UnwrapOr(errorCode.LOGIN_ERROR, "")

	token, _ = util.GenerateToken(util.HmacUser{
		Id:       strconv.Itoa(int(user.ID)),
		Uuid:     user.UUID,
		Username: user.Username,
	})

	return token, user.UUID, user.Username, user.Nickname
}

func (u *UserService) Register(registerDto *dto.RegisterDto) {
	util.Logger().Debug("registerDto", registerDto)
	user := &model.User{Username: registerDto.Username}
	find, _ := user.FindByUsername()
	if find.ID != 0 {
		util.Fail(errorCode.BAD_REQUEST, "用户名已存在")
	}

	salt := util.RandStr(60)
	pwd := util.GenPwd(registerDto.Password, salt)

	user.UUID = uuid.New().String()
	user.Password = pwd
	user.PasswordSalt = salt

	util.NewResult(user.Register()).Unwrap()

	return
}

func (u *UserService) EditUserInfo(userDto *dto.EditUserDto) {
	user := &model.User{ID: userDto.ID}
	user = util.NewResult(user.Find()).UnwrapOr(errorCode.NOT_FOUND_USER, "")
	if userDto.Password != "" {
		util.NewResult(util.CheckPassword(userDto.Password, user.PasswordSalt, user.Password)).Unwrap()
		user.Password = util.GenPwd(userDto.NewPassword, user.PasswordSalt)
	}
	user.Nickname = userDto.Nickname
	user.Email = userDto.Email
	user.Avatar = userDto.Avatar

	util.NewResult(user.Save()).UnwrapOr(errorCode.BAD_REQUEST, "")
}

func (u *UserService) GetUserInfo(uuid string) *model.UserReader {
	user := &model.User{}
	user.UUID = uuid
	userReader := util.NewResult(user.FindByUUID()).UnwrapOr(errorCode.BAD_REQUEST, "请求异常")

	return userReader
}

func (u *UserService) GetUserOrGroupByName(name string) (*model.UserReader, *model.Group) {
	user := &model.User{Username: name}
	userReader := util.NewResult(user.FindByUsername()).UnwrapOrNil().NewUserReader()

	group := &model.Group{Name: name}
	groupRes := util.NewResult(group.FindByName()).UnwrapOrNil()

	return userReader, groupRes
}

func (u *UserService) GetFriendList(id int64, paginatorQo qo.PaginatorQo, searchUserQo qo.SearchUserQo) *paginator.PaginatorCollection[map[string]any] {
	user := &model.User{ID: id}
	res := user.ItemsAndTotal(paginatorQo, searchUserQo)
	paginatorCollection := paginator.NewPaginatorCollection(paginatorQo.Page, res.Items, paginatorQo.PageSize, res.Total)

	return paginatorCollection
}

func (u *UserService) SearchUserList(searchUserQo qo.SearchUserQo) (T *[]map[string]any) {
	user := &model.User{Username: searchUserQo.Username, Nickname: searchUserQo.Nickname, Email: searchUserQo.Email}
	return util.NewResult(user.SearchUserList()).Unwrap()
}

func (u *UserService) AddFriend(addFriendDto *dto.AddFriendDto) {
	util.Logger().Debug("addFriendDto:", addFriendDto)
	// 当前用户
	user := &model.User{ID: addFriendDto.Id}
	user = util.NewResult(user.Find()).Unwrap()
	util.Logger().Debug("user:", user)

	// 添加用户
	friend := &model.User{UUID: addFriendDto.Uuid}
	friendReader := util.NewResult(friend.FindByUUID()).Unwrap()

	userFriend := model.UserFriend{
		UserId:   user.ID,
		FriendId: friendReader.ID,
	}

	if userFriend.Exists() {
		util.Fail(errorCode.BAD_REQUEST, "该用户已经是你好友")
	}

	util.NewResult(userFriend.Save()).Unwrap()
	util.Logger().Debug("userFriend:", userFriend)
}
