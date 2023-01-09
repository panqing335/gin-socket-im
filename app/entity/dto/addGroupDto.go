package dto

type AddGroupDto struct {
	UserId int64  `json:"userId"`
	Name   string `json:"name" form:"name"`
}
