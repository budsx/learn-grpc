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

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Halooo " + req.GetName()}, nil
}

func main() {
	// create connection
	conn, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		grpclog.Fatalln("failed to listen", err.Error())
	}

	// grpc server
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})

	s.Serve(conn)

}
