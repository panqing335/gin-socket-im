package grpc

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"sync"
	"temp/app/grpcService"
	"temp/app/service"
)

var grpcOnce sync.Once

func InitGrpcServer() {
	port := viper.GetString("grpc.port")
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	fmt.Printf("grpc 端口: %v\n", ":"+port)

	s := grpc.NewServer()
	grpcService.RegisterDemoServer(s, &service.DemoServer{})
	reflection.Register(s)

	defer func() {
		s.Stop()
		listen.Close()
	}()

	err = s.Serve(listen)
	if err != nil {
		fmt.Printf("failed to server: %v", err)
		return
	}

}

func Setup() {
	if viper.GetString("grpc.enabled") == "true" {
		grpcOnce.Do(func() {
			go InitGrpcServer()
		})
	}
}
