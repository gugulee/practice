package grpc

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/practice/scratch/grpc/proto"
)


func callUnary(client pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Alice"})
	if err != nil {
		log.Fatalf("SayHello failed: %v", err)
	}
	log.Printf("Unary reply: %s", reply.Message)
}

func callStream(client pb.GreeterClient) {
	stream, err := client.SayHelloStream(context.Background(), &pb.HelloRequest{Name: "Bob"})
	if err != nil {
		log.Fatalf("SayHelloStream failed: %v", err)
	}

	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Stream error: %v", err)
		}
		log.Printf("Stream reply: %s", reply.Message)
	}
}
