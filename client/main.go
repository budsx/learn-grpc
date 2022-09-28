package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "learn-grpc/proto"
)

func main() {
	// Dial
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	c := pb.NewHelloClient(conn)

	req := &pb.HelloRequest{Name: "Budi"}
	res, err := c.SayHello(context.Background(), req)

	if err != nil {
		grpclog.Fatalln(err)
	}

	log.Printf(res.Message)
}
