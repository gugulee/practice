package grpc

import (
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/practice/scratch/grpc/proto"
	"github.com/stretchr/testify/require"
)

func Test_grpc(t *testing.T) {
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		require.NoError(t, err)

		s := grpc.NewServer()
		pb.RegisterGreeterServer(s, &server{})

		log.Println("Server started on :50051")
		err = s.Serve(lis)
		require.NoError(t, err)
	}()

	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	require.NoError(t, err)

	defer conn.Close()
	
	client := pb.NewGreeterClient(conn)

	// Call unary RPC
	callUnary(client)

	// Call streaming RPC
	callStream(client)
}
