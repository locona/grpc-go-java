package main

import (
	"context"
	"log"

	"github.com/locona/grpc-go-java/client-go/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50000"
	// address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	name := defaultName
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: name})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Greet : %v\n", r.Message)
}
