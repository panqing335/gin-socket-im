package model

import (
	"gorm.io/gorm"
	"temp/app/common/mysql"
	"temp/app/entity"
	"temp/app/entity/qo"
	util "temp/app/utils"
	"time"
)

const TableNameUser = "users"

type User struct {
	ID           int64          `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`          // id
	UUID         string         `gorm:"column:uuid;type:varchar(150);not null" json:"uuid"`                  // uuid
	Username     string         `gorm:"column:username;type:varchar(191);not null" json:"username"`          // '用户名'
	Nickname     string         `gorm:"column:nickname;type:varchar(255)" json:"nickname"`                   // 昵称
	Email        string         `gorm:"column:email;type:varchar(80)" json:"email"`                          // 邮箱
	Password     string         `gorm:"column:password;type:varchar(150);not null" json:"password"`          // 密码
	PasswordSalt string         `gorm:"column:password_salt;type:varchar(60);not null" json:"password_salt"` // 密码盐
	Avatar       string         `gorm:"column:avatar;type:varchar(250);not null" json:"avatar"`              // 头像
	CreatedAt    time.Time      `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}

type UserReader struct {
	ID       int64  `json:"id"`
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
}

func (u *User) NewUserReader() (ur *UserReader) {
	if u == nil {
		return
	}
	return &UserReader{
		ID:       u.ID,
		Uuid:     u.UUID,
		Username: u.Username,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		Email:    u.Email,
	}
}

func (u *User) FindByUUID() (user *UserReader, err error) {
	err = mysql.Db.Model(&u).Where("uuid = ?", u.UUID).
		First(&user).Error
	return
}

func (u *User) FindByUsername() (user *User, err error) {
	err = mysql.Db.Model(&u).Where("username = ?", u.Username).
		First(&user).Error
	return
}

func (u *User) FindByNickname() (user *User, err error) {
	err = mysql.Db.Where("nickname = ?", u.Username).
		First(&user).Error
	return
}

func (u *User) FindAllByUsernameAndNickname() (users *[]User, err error) {
	err = mysql.Db.Where("username = ?", u.Username).
		Or("nickname = ?", u.Nickname).
		Find(&users).Error
	return
}

func (u *User) Register() (bool bool, err error) {
	err = mysql.Db.Create(&u).Error
	return
}

func (u *User) Find() (user *User, err error) {
	err = mysql.Db.Where("id = ?", u.ID).Find(&user).Error
	return
}

func (u *User) Save() (user *User, err error) {
	err = mysql.Db.Save(&u).Error
	return
}

//type ItemsAndTotal struct {
//	Items *[]map[string]any
//	Total int64
//}

func (u *User) ItemsAndTotal(paginatorQo qo.PaginatorQo, searchUserQo qo.SearchUserQo) (res *entity.ItemsAndTotal) {
	var items *[]struct {
		UserReader `gorm:"embedded"`
	}
	var total int64

	//mysql.Db.Table(u.TableName()+" u"). // 无 SoftDeletes
	mysql.Db.Model(&User{}).
		Joins("JOIN user_friends AS uf ON uf.friend_id = users.id").
		Select("users.id, users.username, users.uuid, users.avatar, users.nickname").
		Where("uf.user_id = ?", u.ID).
		Where(&searchUserQo).
		Count(&total). // sql1: SELECT count(*) FROM `users` JOIN user_friends AS uf ON uf.friend_id = users.id WHERE uf.user_id = 1 AND `users`.`deleted_at` IS NULL
		Offset((paginatorQo.Page - 1) * paginatorQo.PageSize).
		Limit(paginatorQo.PageSize).
		Scan(&items) // sql2: SELECT users.id, users.username, users.uuid, users.avatar, users.nickname FROM `users` JOIN user_friends AS uf ON uf.friend_id = users.id WHERE uf.user_id = 1 AND `users`.`deleted_at` IS NULL LIMIT 10 OFFSET 10

	// 转换成[]map[string]any 方便统一转换
	c := entity.ItemsAndTotal{}
	c.Total = total
	c.Items = util.MapToStruct[[]map[string]any](items)

	return &c
}

func (u *User) SearchUserList() (*[]map[string]any, error) {
	var items *[]struct {
		UserReader `gorm:"embedded"`
	}
	err := mysql.Db.Model(&User{}).
		Where("nickname = ?", u.Nickname).
		Or("username = ?", u.Username).
		Or("email = ?", u.Email).
		Scan(&items).
		Error
	return util.MapToStruct[[]map[string]any](items), err
}
