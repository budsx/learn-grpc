# GRPC Golang

## Proto

Protocol Buffers adalah metode untuk serialisasi data terstruktur, yang dibuat oleh Google. Protobuf cocok digunakan pada aplikasi yang berkomunikasi dengan aplikasi lain.

gRPC adalah sebuah remote procedure call atau RPC yang dibuat oleh google. gRPC menggunakan HTTP/2 untuk komunikasinya, dan Protocol Buffers di bagian antarmuka-nya. 

Rest api payloadnya menggunakan JSON sementara gRPC menggunakan protobuf, komunikasi dengan grpc artinya 

```
syntax = "proto3";

```
Menggunakan syntax versi proto3.

```
package hello;

option go_package = "./hello";

```
package hello -> nama package setelah dicompile.
go_package -> tempat hasil compile proto ke golang file.

```
message HelloRequest {
    string name = 1;
}

message HelloResponse {
    string message = 1;
}

```
message -> model atau dao 

```
service Hello {
    rpc SayHello(HelloRequest) returns (HelloResponse) {}
}
```

service Hello -> definisi method sayHello dengan membutuhkan HelloRequest dan return HelloResponse

Compile
```
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
proto/hello.proto
```

Hasil compile akan menjadi file hello_grpc.pb.go dan hello.pb.go



## Server
```
type server struct {
	pb.UnimplementedHelloServer
}
```
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Halooo " + req.GetName()}, nil
}
```
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Halooo " + req.GetName()}, nil
}
```
Implement HelloServer dan method SayHello yang terdapat di hello_grpc.pb.go

```
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})

	s.Serve(conn)
```
Create new grpc server and serve.


## Client

```
conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
```
Client melakukan DIAL (memanggil server grpc)

```
c := pb.NewHelloClient(conn)
```
NewHelloClient adalah constructor yang ada di hello_grpc.pb.go


```
req := &pb.HelloRequest{Name: "Budi"}
res, err := c.SayHello(context.Background(), req)    
````
Client mengirim request (HellRequest) dan memanggilnya dalam fungsi SayHello()


