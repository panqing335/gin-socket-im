package model

import (
	"fmt"
	"gorm.io/gorm"
	"temp/app/common/cache"
	"temp/app/common/mysql"
	"temp/app/constants/errorCode"
	"temp/app/entity/qo"
	util "temp/app/utils"
	"time"
)

const TableNameTempUser = "temp_user"

// TempUser mapped from table <temp_user>
type TempUser struct {
	ID           int64          `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	Username     string         `gorm:"column:username;type:varchar(128);not null" json:"username"`
	Password     string         `gorm:"column:password;type:varchar(128);not null" json:"password"`
	PasswordSalt string         `gorm:"column:password_salt;type:varchar(60);not null" json:"password_salt"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
}

func NewTempUser(ID int64, username string, password string, passwordSalt string, createdAt time.Time, updatedAt time.Time, deletedAt gorm.DeletedAt) *TempUser {
	return &TempUser{ID: ID, Username: username, Password: password, PasswordSalt: passwordSalt, CreatedAt: createdAt, UpdatedAt: updatedAt, DeletedAt: deletedAt}
}

// TableName TempUser's table name
func (*TempUser) TableName() string {
	return TableNameTempUser
}

func (t *TempUser) AfterSave(*gorm.DB) (err error) {
	fmt.Println("after: save")
	fromCache := cache.NewFindFromCache(t, err)
	fromCache.DelCache(t.TableName(), t.ID)
	return
}

func (t *TempUser) AfterDelete(*gorm.DB) (err error) {
	fmt.Println("after: delete")
	fromCache := cache.NewFindFromCache(t, err)
	fromCache.DelCache(t.TableName(), t.ID)
	return
}

func (t *TempUser) Find() (tempUser *TempUser, err error) {
	result := mysql.Db.First(&tempUser, t.ID)
	if result.Error != nil {
		err = result.Error
		return
	}

	return
}

func (t *TempUser) FindByUsername() (tempUser *TempUser, err error) {
	result := mysql.Db.Where("username = ?", t.Username).
		Select([]string{"id", "username", "password", "password_salt"}).
		First(&tempUser)

	if result.Error != nil {
		err = result.Error
	}

	return
}

func (t *TempUser) Create() (bool, err error) {
	result := mysql.Db.Create(&t)
	if result.Error != nil {
		err = result.Error
	}
	return
}

func (t *TempUser) Update() (bool, err error) {
	result := mysql.Db.Updates(&t)
	if result.Error != nil {
		err = result.Error
		util.Fail(errorCode.BAD_REQUEST, "")
	}
	return
}

func (t *TempUser) Destroy() (bool, err error) {
	result := mysql.Db.Delete(&t)
	if result.Error != nil {
		err = result.Error
		util.Fail(errorCode.BAD_REQUEST, "")
	}
	return
}

func (t *TempUser) FindFromCache() (ret map[string]string, err error) {
	fromCache := cache.NewFindFromCache(t, err)
	ret, err = fromCache.FindFromCache(t.TableName(), t.ID)
	return
}

type Paginator struct {
	CurrentPage int
	Items       []TempUser
	PageSize    int
	Total       int64
}

func NewPaginator(currentPage int, items []TempUser, pageSize int, total int64) *Paginator {
	return &Paginator{CurrentPage: currentPage, Items: items, PageSize: pageSize, Total: total}
}

func (t *TempUser) Paginator(paginatorQo qo.PaginatorQo, searchQo qo.SearchTempUserQo) (paginator *Paginator, err error) {
	var items *[]TempUser
	mysql.Db.Table(t.TableName()).Where(&searchQo).Offset((paginatorQo.Page - 1) * paginatorQo.PageSize).Limit(paginatorQo.PageSize).Find(&items)
	var total int64
	mysql.Db.Count(&total)

	paginator = NewPaginator(paginatorQo.Page, *items, paginatorQo.Page, total)
	return
}

func (t *TempUser) ItemsAndTotal(paginatorQo qo.PaginatorQo, searchQo qo.SearchTempUserQo) (items *[]TempUser, total int64) {
	mysql.Db.Table(t.TableName()).Where(&searchQo).Offset((paginatorQo.Page - 1) * paginatorQo.PageSize).Limit(paginatorQo.PageSize).Find(&items)
	mysql.Db.Table(t.TableName()).Where(&searchQo).Count(&total)
	return
}
