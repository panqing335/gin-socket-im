package entity

import "time"

type MsgResponse struct {
	ID           int64     `json:"id" gorm:"primarykey"`
	FromUserId   int64     `json:"fromUserId" gorm:"index"`
	ToUserId     int64     `json:"toUserId" gorm:"index"`
	Content      string    `json:"content" gorm:"type:varchar(2500)"`
	ContentType  int64     `json:"contentType" gorm:"comment:'消息内容类型：1文字，2语音，3视频'"`
	CreatedAt    time.Time `json:"createAt"`
	FromUsername string    `json:"fromUsername"`
	ToUsername   string    `json:"toUsername"`
	Avatar       string    `json:"avatar"`
	Url          string    `json:"url"`
}
