package service

import (
	"context"
	"fmt"
	"temp/app/grpcService"
)

type DemoServer struct {
	grpcService.UnimplementedDemoServer
}

func (s *DemoServer) UnaryCall(ctx context.Context, req *grpcService.DemoRequest) (*grpcService.DemoReply, error) {
	fmt.Println("request:", req.Json)

	return &grpcService.DemoReply{Message: "Hello, world!" + req.Json}, nil
}
