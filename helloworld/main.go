package main

import (
	"context"
	"log"
	"net"
	"strings"

	"github.com/shugo256/nextjs-grpc/helloworld/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type HelloWorldService struct{
	proto.UnimplementedHelloWorldServiceServer
}

func (s HelloWorldService) SayHello(_ context.Context, req *proto.HelloWorldRequest) (resp *proto.HelloWorldResponse, err error) {
	log.Println(req.String())
	resp = &proto.HelloWorldResponse{
		Message: "Hello World" + strings.Repeat("!", int(req.NumRepeat)),
	}
	return
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("HelloWorldService started at localhost:8000!")
	
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	proto.RegisterHelloWorldServiceServer(grpcServer, HelloWorldService{})
	grpcServer.Serve(lis)
}
