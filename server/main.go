package main

import (
	"context"
	"net"

	pb "learn-grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type server struct {
	pb.UnimplementedHelloServer
}



func (h *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// create connection
	conn, err := net.Listen("tcp", ":8080")
	if err != nil {
		grpclog.Fatalln("failed to listen", err.Error())
	}

	// grpc server
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})

	s.Serve(conn)

}
