package server

import (
	"context"
	"log"
	"net"

	pb "github.com/guilhermelinosp/hello-grpc"
	"google.golang.org/grpc"
)

type server struct {
 pb.UnimplementedHelloWorldServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
 return &pb.HelloWorldResponse{Message: "Hello, World! "}, nil
}

func main() {
 lis, err := net.Listen("tcp", ":50051")
 if err != nil {
  log.Fatalf("failed to listen on port 50051: %v", err)
 }

 s := grpc.NewServer()
 pb.RegisterHelloWorldServiceServer(s, &server{})
 log.Printf("gRPC server listening at %v", lis.Addr())
 if err := s.Serve(lis); err != nil {
  log.Fatalf("failed to serve: %v", err)
 }
}