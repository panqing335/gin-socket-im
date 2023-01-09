package cache

import (
	"log"
	"temp/app/common/mysql"
	"temp/app/common/redis"
)

func DbGetter(tableName string, id int64) DbGetFunc {
	return func() redis.ResultType {
		log.Print("from db")
		var result map[string]any
		mysql.Db.Table(tableName).Where("id = ?", id).Find(&result)
		return result
	}
}
