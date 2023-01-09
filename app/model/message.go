package model

import (
	"gorm.io/gorm"
	"temp/app/common/mysql"
	"temp/app/entity"
	"time"
)

const TableNameMessage = "messages"

type Message struct {
	ID          int64          `json:"id" gorm:"primarykey"`
	FromUserId  int64          `json:"fromUserId" gorm:"index"`
	ToUserId    int64          `json:"toUserId" gorm:"index;comment:'发送给端的id，可为用户id或者群id'"`
	Content     string         `json:"content" gorm:"type:varchar(2500)"`
	MessageType int16          `json:"messageType" gorm:"comment:'消息类型：1单聊，2群聊'"`
	ContentType int16          `json:"contentType" gorm:"comment:'消息内容类型：1文字 2.普通文件 3.图片 4.音频 5.视频 6.语音聊天 7.视频聊天'"`
	Pic         string         `json:"pic" gorm:"type:text;comment:'缩略图"`
	Url         string         `json:"url" gorm:"type:varchar(350);comment:'文件或者图片地址'"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
}

func NewMessage(ID int64, fromUserId int64, toUserId int64, content string, messageType int16, contentType int16, pic string, url string) *Message {
	return &Message{ID: ID, FromUserId: fromUserId, ToUserId: toUserId, Content: content, MessageType: messageType, ContentType: contentType, Pic: pic, Url: url}
}

// TableName User's table name
func (*Message) TableName() string {
	return TableNameMessage
}

func (m *Message) Save() {
	mysql.Db.Save(&m)
}

func (m *Message) ScanUser(userId int64, friendId int64) (message []entity.MsgResponse, err error) {
	err = mysql.Db.Raw("SELECT m.id, m.from_user_id, m.to_user_id, m.content, m.content_type, m.url, "+
		"m.created_at, u.username AS from_username, u.avatar, to_user.username AS to_username  "+
		"FROM messages AS m LEFT JOIN users AS u ON m.from_user_id = u.id "+
		"LEFT JOIN users AS to_user ON m.to_user_id = to_user.id "+
		"WHERE message_type = 1 and from_user_id IN (?, ?) AND to_user_id IN (?, ?)",
		userId, friendId, userId, friendId).Scan(&message).Error
	return
}

func (m *Message) ScanGroup(groupId int64) (message []entity.MsgResponse, err error) {
	err = mysql.Db.Raw("SELECT m.id, m.from_user_id, m.to_user_id, m.content, m.content_type, m.url, m.created_at, "+
		"u.username AS from_username, u.avatar "+
		"FROM messages AS m "+
		"LEFT JOIN users AS u ON m.from_user_id = u.id "+
		"WHERE m.message_type = 2 AND m.to_user_id = ?",
		groupId).Scan(&message).Error
	return
}
