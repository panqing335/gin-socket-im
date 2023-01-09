package socket

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"temp/app/constants/contentType"
	"temp/app/grpcService"
	"temp/app/kafka"
	util "temp/app/utils"
)

type Client struct {
	Conn *websocket.Conn
	Name string
	Send chan []byte
}

func (c *Client) Close() {
	err := c.Conn.Close()
	if err != nil {
		return
	}
}

func (c *Client) Read() {
	defer func() {
		MyServer.UnRegister <- c
		c.Close()
	}()

	for {
		c.Conn.PongHandler()
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			util.Logger().Error("client read message error:", err.Error())
			MyServer.UnRegister <- c
			c.Close()
			break
		}
		msg := &grpcService.Message{}
		proto.Unmarshal(p, msg)

		if msg.Type == contentType.HEAT_BEAT {
			pong := &grpcService.Message{
				Content: contentType.PONG,
				Type:    contentType.HEAT_BEAT,
			}
			pongByte := util.NewResult(proto.Marshal(pong)).UnwrapOrNil()
			fmt.Printf("pongByte %v\n", pongByte)
			c.Conn.WriteMessage(websocket.BinaryMessage, pongByte)
		} else {
			if viper.GetString("msgChannel.type") == contentType.KAFKA {
				kafka.Send(p)
			} else {
				MyServer.Broadcast <- p
			}
		}
	}
}

func (c *Client) Write() {
	defer func() {
		c.Close()
	}()

	for message := range c.Send {
		c.Conn.WriteMessage(websocket.BinaryMessage, message)
	}
}
