package service

import (
	"temp/app/constants/contentType"
	"temp/app/constants/errorCode"
	"temp/app/entity"
	"temp/app/entity/dto"
	"temp/app/grpcService"
	"temp/app/model"
	"temp/app/utils"
)

type MessageService struct{}

func NewMessageService() *MessageService {
	return &MessageService{}
}

func (m *MessageService) GetMessage(msgReqDto dto.MsgReqDto) (messageArr []entity.MsgResponse) {
	if msgReqDto.MessageType == contentType.MESSAGE_TYPE_USER {
		user := &model.User{ID: msgReqDto.ID}
		userReader := util.NewResult(user.Find()).UnwrapOr(errorCode.BAD_REQUEST, "获取信息失败，请重新尝试")

		friend := &model.User{UUID: msgReqDto.FriendUuid}
		friendReader := util.NewResult(friend.FindByUUID()).UnwrapOr(errorCode.BAD_REQUEST, "获取信息失败，请重新尝试")

		message := &model.Message{}
		messageArr = util.NewResult(message.ScanUser(userReader.ID, friendReader.ID)).UnwrapOrNil()
		return
	}

	if msgReqDto.MessageType == contentType.MESSAGE_TYPE_GROUP {
		messageArr = util.NewResult(fetchGroupMsg(msgReqDto.FriendUuid)).Unwrap()
		return
	}

	return nil
}

func fetchGroupMsg(toUuid string) (messageArr []entity.MsgResponse, err error) {
	group := &model.Group{UUID: toUuid}
	group = util.NewResult(group.FindByUUID()).Unwrap()

	message := model.Message{}
	messageArr = util.NewResult(message.ScanGroup(group.ID)).Unwrap()
	return
}

func (m *MessageService) SaveMessage(message grpcService.Message) {
	user := &model.User{
		UUID: message.From,
	}
	userReader := util.NewResult(user.FindByUUID()).Unwrap()

	var toUserId int64 = 0

	if message.MessageType == contentType.MESSAGE_TYPE_USER {
		toUser := &model.User{
			UUID: message.To,
		}
		toUserReader := util.NewResult(toUser.FindByUUID()).Unwrap()
		toUserId = toUserReader.ID
	}

	if message.MessageType == contentType.MESSAGE_TYPE_GROUP {
		group := &model.Group{UUID: message.To}
		reader := util.NewResult(group.FindByUUID()).Unwrap()
		toUserId = reader.ID
	}

	saveMessage := model.Message{
		FromUserId:  userReader.ID,
		ToUserId:    toUserId,
		Content:     message.Content,
		MessageType: int16(message.MessageType),
		ContentType: int16(message.ContentType),
		Url:         message.Url,
	}

	saveMessage.Save()
}
