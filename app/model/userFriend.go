package model

import (
	"gorm.io/gorm"
	"temp/app/common/mysql"
	"time"
)

type UserFriend struct {
	ID        int64          `json:"id" gorm:"primarykey"`
	UserId    int64          `json:"userId" gorm:"index;comment:'用户ID'"`
	FriendId  int64          `json:"friendId" gorm:"index;comment:'好友ID'"`
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
}

func NewUserFriend(ID int64, userId int64, friendId int64) *UserFriend {
	return &UserFriend{ID: ID, UserId: userId, FriendId: friendId}
}

func (f *UserFriend) Exists() bool {
	var count int64
	mysql.Db.Where("user_id=?", f.UserId).Where("friend_id=?", f.FriendId).Count(&count)
	return count > 0
}

func (f *UserFriend) Save() (uf *UserFriend, err error) {
	err = mysql.Db.Save(&f).Error
	return
}
