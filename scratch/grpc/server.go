package grpc

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/practice/scratch/grpc/proto"
)

type server struct {
	pb.UnimplementedGreeterServer
}

// Unary RPC implementation
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received unary request: %v", req.Name)
	return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}

// Server streaming RPC implementation
func (s *server) SayHelloStream(req *pb.HelloRequest, stream pb.Greeter_SayHelloStreamServer) error {
	log.Printf("Received stream request: %v", req.Name)

	for i := 0; i < 5; i++ {
		reply := &pb.HelloReply{
			Message: fmt.Sprintf("Hello %s (#%d)", req.Name, i+1),
		}
		if err := stream.Send(reply); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
