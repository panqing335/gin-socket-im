package dto

type MsgReqDto struct {
	ID          int64  `json:"id"`
	MessageType int64  `json:"messageType" form:"MessageType"`
	FriendUuid  string `json:"friendUuid" form:"FriendUuid"`
}
