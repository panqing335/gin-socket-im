package socket

import (
	"encoding/base64"
	"github.com/gogo/protobuf/proto"
	uuid "github.com/satori/go.uuid"
	"os"
	"strings"
	"sync"
	"temp/app/constants/contentType"
	"temp/app/grpcService"
	"temp/app/service"
	util "temp/app/utils"
)

var MyServer = NewServer()

type Server struct {
	Clients    map[string]*Client
	mutex      *sync.Mutex
	Broadcast  chan []byte
	Register   chan *Client
	UnRegister chan *Client
}

func NewServer() *Server {
	return &Server{
		Clients:    make(map[string]*Client),
		mutex:      &sync.Mutex{},
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
	}
}

// ConsumerKafkaMsg 消费kafka里面的消息, 然后直接放入go channel中统一进行消费 /**
func ConsumerKafkaMsg(data []byte) {
	MyServer.Broadcast <- data
}

func (s *Server) Start() {
	util.Logger().Info("socket start")
	for {
		select {
		case conn := <-s.Register:
			util.Logger().Info("login in " + conn.Name)
			s.Clients[conn.Name] = conn
			msg := &grpcService.Message{
				From:    "System",
				To:      conn.Name,
				Content: "welcome!",
			}
			marshal, _ := proto.Marshal(msg)
			conn.Send <- marshal
		case conn := <-s.UnRegister:
			util.Logger().Info("logout " + conn.Name)
			if _, ok := s.Clients[conn.Name]; ok {
				close(conn.Send)
				delete(s.Clients, conn.Name)
			}
		case message := <-s.Broadcast:
			util.Logger().Info("broadcast " + string(message))
			msg := &grpcService.Message{}
			err := proto.Unmarshal(message, msg)
			if err != nil {
				return
			}
			if msg.To != "" {
				if msg.ContentType >= contentType.TEXT && msg.ContentType <= contentType.VIDEO {
					_, exits := s.Clients[msg.From]
					if exits {
						SaveMessage(msg)
					}

					if msg.MessageType == contentType.MESSAGE_TYPE_USER {
						client, ok := s.Clients[msg.To]
						if ok {
							marshal, err := proto.Marshal(msg)
							if err == nil {
								client.Send <- marshal
							}
						}
					} else if msg.MessageType == contentType.MESSAGE_TYPE_GROUP {
						s.SendGroupMsg(msg)
					}
				} else {
					client, ok := s.Clients[msg.To]
					if ok {
						client.Send <- message
					}
				}
			} else {
				// 无对应接受人员进行广播
				for id, conn := range s.Clients {
					util.Logger().Info("allUser id:" + id)

					select {
					case conn.Send <- message:
					default:
						close(conn.Send)
						delete(s.Clients, conn.Name)
					}
				}
			}
		}
	}
}

func (s *Server) SendGroupMsg(msg *grpcService.Message) {
	groupService := service.GroupService{}
	users := groupService.GetUsersByGroupUuid(msg.To)

	for _, user := range *users {
		if user.Uuid == msg.From {
			continue
		}

		client, ok := s.Clients[user.Uuid]
		if !ok {
			continue
		}

		userService := service.UserService{}
		fromUserInfos := userService.GetUserInfo(msg.From)

		msgSend := &grpcService.Message{
			Avatar:       fromUserInfos.Avatar,
			FromUsername: msg.FromUsername,
			From:         msg.To,
			To:           msg.From,
			Content:      msg.Content,
			ContentType:  msg.ContentType,
			Type:         msg.Type,
			MessageType:  msg.MessageType,
		}

		msgBytes, err := proto.Marshal(msgSend)
		if err == nil {
			client.Send <- msgBytes
		}
	}
}

// SaveMessage 保存消息
func SaveMessage(msg *grpcService.Message) {
	if msg.ContentType == contentType.FILE {
		url := uuid.NewV1().String() + ".png"
		index := strings.Index(msg.Content, "base64")
		index += 7

		content := msg.Content
		content = content[index:]

		dataBuffer, err := base64.StdEncoding.DecodeString(content)
		if err != nil {
			util.Logger().Error("transfer base64 to file error", err.Error())
			return
		}

		err = os.WriteFile(url, dataBuffer, 0666)
		if err != nil {
			util.Logger().Error("write file error", err.Error())
			return
		}
		msg.Url = url
		msg.Content = ""
	} else if msg.ContentType == contentType.IMAGE {
		// 普通的文件二进制上传
		fileSuffix := util.GetFileType(msg.File)
		nullStr := ""
		if nullStr == fileSuffix {
			fileSuffix = strings.ToLower(msg.FileSuffix)
		}
		contentTypeBySuffix := util.GetContentTypeBySuffix(fileSuffix)
		url := uuid.NewV1().String() + "." + fileSuffix
		err := os.WriteFile(url, msg.File, 0666)
		if err != nil {
			util.Logger().Error("write file error", err.Error())
			return
		}
		msg.Url = url
		msg.File = nil
		msg.ContentType = contentTypeBySuffix
	}

	messageService := service.NewMessageService
	messageService().SaveMessage(*msg)
}

var socketOnce sync.Once

func Setup() {
	go socketOnce.Do(func() {
		MyServer.Start()
	})
}
