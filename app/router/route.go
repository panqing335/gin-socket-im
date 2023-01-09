package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"temp/app/controller/group"
	"temp/app/controller/msg"
	"temp/app/controller/user"
	"temp/app/middleware"
	util "temp/app/utils"
	_ "temp/config"
	"time"
)

var router *gin.Engine

func Setup() {
	router = gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.Recover)
	router.Use(middleware.Cors)
	GroupDefault()
	GroupUser()
	Group()
	GroupFile()
	GroupSocket()

	port := viper.GetString("server.port")
	fmt.Println("当前端口：", port)
	s := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		util.Logger().Error("server error: ", err)
	}
}

func GroupDefault() {
	router.POST("/login", user.Login)
	router.POST("/register", user.Register)
}

func GroupUser() *gin.RouterGroup {
	v1 := router.Group("user")
	v1.Use(middleware.Auth)
	{
		v1.GET("/getFriendList", user.GetFriendList)
		v1.POST("/editUserInfo", user.EditUserInfo)
		v1.GET("/getUserInfo", user.GetUserInfo)
		v1.GET("/getUserOrGroupByName", user.GetUserOrGroupByName)
		v1.GET("/searchUserList", user.SearchUserList)
		v1.POST("/addFriend", user.AddFriend)
		v1.GET("/msg", msg.GetMsg)
	}
	return v1
}

func Group() *gin.RouterGroup {
	g := router.Group("group")
	g.Use(middleware.Auth)
	{
		g.GET("/paginator", group.GetGroups)
		g.POST("/add", group.AddGroup)
		g.POST("/join", group.JoinGroup)
	}
	return g
}

func GroupFile() *gin.RouterGroup {
	f := router.Group("")
	//f.Use(middleware.Auth)
	{
		f.GET("/file/:fileName", msg.GetFile)
		f.POST("/file", msg.SaveFile)
	}

	return f
}

func GroupSocket() *gin.RouterGroup {
	socket := RunSocket
	ws := router.Group("")
	ws.Use(middleware.SocketAuth)
	{
		ws.GET("/socket.io", socket)
	}

	return ws
}
