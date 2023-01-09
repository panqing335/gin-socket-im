package mysql

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"net/url"
	"sync"
	util "temp/app/utils"
	"time"
)

var Db *gorm.DB

type MyWriter struct {
	mlog *logrus.Logger
}

// 实现gorm/logger.Writer接口
func (m *MyWriter) Printf(format string, v ...interface{}) {
	logstr := fmt.Sprintf(format, v...)
	//利用loggus记录日志
	m.mlog.Info(logstr)
}

func NewMyWriter() *MyWriter {
	return &MyWriter{mlog: util.Logger()}
}

var mysqlOnce sync.Once

func Setup() {
	mysqlOnce.Do(func() {
		host := viper.GetString("datasource.host")
		port := viper.GetString("datasource.port")
		database := viper.GetString("datasource.database")
		username := viper.GetString("datasource.username")
		password := viper.GetString("datasource.password")
		charset := viper.GetString("datasource.charset")
		loc := viper.GetString("datasource.loc")

		sqlStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
			username,
			password,
			host,
			port,
			database,
			charset,
			url.QueryEscape(loc),
		)
		fmt.Println("数据库连接:", sqlStr)

		// 配置日志输出
		newLogger := logger.New(
			NewMyWriter(),
			logger.Config{
				//慢SQL阈值
				SlowThreshold: time.Millisecond,
				//设置日志级别，只有Warn以上才会打印sql
				LogLevel: logger.Info,
			},
		)

		gdb, err := gorm.Open(mysql.Open(sqlStr), &gorm.Config{
			Logger: newLogger,
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: viper.GetString("datasource.tablePrefix"),
				//SingularTable: true,
				//NameReplacer:  string.NewReplacer(),
				//NoLowerCase:   true,
			},
		})

		if err != nil {
			fmt.Println("打开数据库失败", err)
			panic("打开数据库失败" + err.Error())
		}
		err = gdb.Use(
			dbresolver.Register(dbresolver.Config{}).
				SetMaxOpenConns(100).
				SetMaxIdleConns(20),
		)
		if err != nil {
			return
		}

		Db = gdb
	})
}
