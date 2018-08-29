package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/locona/grpc-go-java/server-go/proto"
	"google.golang.org/grpc"
)

const (
	port = 50000
)

type greeterServer struct{}

func (g *greeterServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: fmt.Sprintf("Yaa!! %v", req.Name)}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal(err)
	}
	opts := make([]grpc.ServerOption, 0)
	log.Println("-------> START")

	grpcServer := grpc.NewServer(opts...)
	proto.RegisterGreeterServer(grpcServer, &greeterServer{})
	grpcServer.Serve(lis)
}
