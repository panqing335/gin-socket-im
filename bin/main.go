package main

import (
	Cron "temp/app/common/cron"
	Grpc "temp/app/common/grpc"
	MQ "temp/app/common/mq"
	Mysql "temp/app/common/mysql"
	Redis "temp/app/common/redis"
	Socket "temp/app/common/socket"
	Router "temp/app/router"
	util "temp/app/utils"
	Config "temp/config"
)

func main() {
	Config.Setup()
	Redis.Setup()
	Mysql.Setup()
	Grpc.Setup()
	Cron.Setup()
	MQ.Setup()
	Socket.Setup()
	Router.Setup()
}

func genModel() {
	var db = Mysql.Db
	util.TestGEN(db)
}
