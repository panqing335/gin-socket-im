package room

type JoinRoomDto struct {
	RoomUuid string `json:"roomUuid" form:"roomUuid"`
	Pwd      string `json:"pwd" form:"pwd"`
	UserId   int64
}
