package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"temp/app/common/socket"
	util "temp/app/utils"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RunSocket(c *gin.Context) {
	uuid := util.GetUUID(c)
	if uuid == "" {
		return
	}
	hh := http.Header{}
	hh.Set("Sec-Websocket-Protocol", c.GetHeader("Sec-Websocket-Protocol"))
	ws := util.NewResult(upGrader.Upgrade(c.Writer, c.Request, hh)).Unwrap()

	client := &socket.Client{
		Name: uuid,
		Conn: ws,
		Send: make(chan []byte),
	}

	socket.MyServer.Register <- client

	go client.Read()
	go client.Write()
}
