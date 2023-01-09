package cache

import (
	"fmt"
	"github.com/spf13/viper"
	"strconv"
	"temp/app/constants/errorCode"
	util "temp/app/utils"
)

type Models interface {
	map[string]any | map[string]string | any
}

type FindFromCache[T Models] struct {
	data T
	err  error
}

func NewFindFromCache[T Models](data T, err error) *FindFromCache[T] {
	return &FindFromCache[T]{data: data, err: err}
}

func (f *FindFromCache[T]) FindFromCache(tableName string, id int64) (ret map[string]string, err error) {
	simpleCache := Cache()
	defer ReleaseCache(simpleCache)

	simpleCache.DbGetter = DbGetter(tableName, id)

	keyString := viper.Get("modelCache.key")
	cacheKey := fmt.Sprintf(keyString.(string)+"\n", tableName, "id", strconv.FormatInt(id, 10))

	ret = simpleCache.GetCache(cacheKey)
	if ret["id"] == "-1" {
		util.Fail(errorCode.BAD_REQUEST, "ID 不存在")
	}
	return
}

func (f *FindFromCache[T]) DelCache(tableName string, id int64) {
	simpleCache := Cache()
	defer ReleaseCache(simpleCache)

	keyString := viper.Get("modelCache.key")
	cacheKey := fmt.Sprintf(keyString.(string)+"\n", tableName, "id", strconv.FormatInt(id, 10))

	simpleCache.DelCache(cacheKey)
}
