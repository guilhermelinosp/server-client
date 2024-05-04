package client

import (
	"context"
	"log"
	"time"

	pb "github.com/guilhermelinosp/hello-grpc/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
 conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
 if err != nil {
  log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
 }
 defer conn.Close()
 c := pb.NewHelloClient(conn)

 ctx, cancel := context.WithTimeout(context.Background(), time.Second)
 defer cancel()

 r, err := c.SayHello(ctx, &pb.HelloRequest{})
 if err != nil {
  log.Fatalf("error calling function SayHello: %v", err)
 }

 log.Printf("Response from gRPC server's SayHello function: %s", r.GetOutput())
}