package model

import (
	"gorm.io/gorm"
	"temp/app/common/mysql"
	"time"
)

type GroupMember struct {
	ID        int64          `json:"id" gorm:"primarykey"`
	UserId    int64          `json:"userId" gorm:"index;comment:'用户ID'"`
	GroupId   int64          `json:"groupId" gorm:"index;comment:'群组ID'"`
	Nickname  string         `json:"nickname" gorm:"type:varchar(350);comment:'昵称"`
	Mute      int16          `json:"mute" gorm:"comment:'是否禁言'"`
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
}

func (gm *GroupMember) Save() (groupMember *GroupMember, err error) {
	err = mysql.Db.Save(&gm).Error
	return
}

func (gm *GroupMember) Exists() bool {
	var count int64
	mysql.Db.Model(&gm).Where("user_id = ? AND group_id = ?", gm.UserId, gm.GroupId).Count(&count)
	if count == 0 {
		return false
	}
	return true
}
