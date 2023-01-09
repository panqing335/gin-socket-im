package dto

type AddFriendDto struct {
	Id   int64  `json:"id"`
	Uuid string `json:"uuid" from:"uuid"`
}
