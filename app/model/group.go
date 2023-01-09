package model

import (
	"gorm.io/gorm"
	"temp/app/common/mysql"
	"temp/app/entity"
	"temp/app/entity/qo"
	util "temp/app/utils"
	"time"
)

const TableNameGroup = "groups"

type Group struct {
	ID        int64          `json:"id" gorm:"primarykey"`
	UUID      string         `json:"uuid" gorm:"type:varchar(150);not null;unique_index:idx_uuid;comment:'uuid'"`
	UserId    int64          `json:"userId" gorm:"index;comment:'群主ID'"`
	Name      string         `json:"name" gorm:"type:varchar(150);comment:'群名称"`
	Notice    string         `json:"notice" gorm:"type:varchar(350);comment:'群公告"`
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
}

type GroupReader struct {
	ID        int64     `json:"groupId"`
	UUID      string    `json:"uuid" gorm:"type:varchar(150);not null;unique_index:idx_uuid;comment:'uuid'"`
	UserId    int64     `json:"userId" gorm:"index;comment:'群主ID'"`
	Name      string    `json:"name" gorm:"type:varchar(150);comment:'群名称"`
	Notice    string    `json:"notice" gorm:"type:varchar(350);comment:'群公告"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
}

// TableName User's table name
func (*Group) TableName() string {
	return TableNameGroup
}

func (g *Group) FindByUUID() (group *Group, err error) {
	err = mysql.Db.Where("uuid = ?", g.UUID).
		First(&group).Error

	return
}

func (g *Group) FindByName() (group *Group, err error) {
	err = mysql.Db.Model(&g).Where("name = ?", g.Name).
		First(&group).Error
	return
}

func (g *Group) ItemsAndTotal(paginatorQo qo.PaginatorQo) (res *entity.ItemsAndTotal) {
	var items *[]struct {
		GroupReader `gorm:"embedded"`
	}
	var total int64

	mysql.Db.Model(&Group{}).
		Joins("JOIN group_members AS gm ON gm.group_id = groups.id").
		Select("groups.id, groups.uuid, groups.created_at, groups.name, groups.notice").
		Where("gm.user_id = ?", g.UserId).
		Count(&total). // sql1: SELECT count(*) FROM `users` JOIN user_friends AS uf ON uf.friend_id = users.id WHERE uf.user_id = 1 AND `users`.`deleted_at` IS NULL
		Offset((paginatorQo.Page - 1) * paginatorQo.PageSize).
		Limit(paginatorQo.PageSize).
		Order("id DESC").
		Scan(&items) // sql2: SELECT users.id, users.username, users.uuid, users.avatar, users.nickname FROM `users` JOIN user_friends AS uf ON uf.friend_id = users.id WHERE uf.user_id = 1 AND `users`.`deleted_at` IS NULL LIMIT 10 OFFSET 10

	// 转换成[]map[string]any 方便统一转换
	c := entity.ItemsAndTotal{}
	c.Total = total
	c.Items = util.MapToStruct[[]map[string]any](items)

	return &c
}

func (g *Group) Save() (group *Group, err error) {
	err = mysql.Db.Save(&g).Error
	return g, err
}

func (g *Group) FindUsersByGroupId() (*[]UserReader, error) {
	var users *[]UserReader
	err := mysql.Db.Model(&g).
		Joins("JOIN group_members AS gm ON gm.group_id = groups.id").
		Joins("JOIN users As u ON u.id = gm.user_id").
		Where("groups.id = ?", g.ID).
		Select("u.uuid, u.avatar, u.username").
		Scan(&users).Error

	return users, err
}
