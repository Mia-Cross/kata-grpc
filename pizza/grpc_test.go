package pizza

import (
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
)

func initGRPCServer() *grpc.Server {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	return grpcServer
}

func TestGetPizzaFromName(t testing.T) {
	grpcServer := initGRPCServer()

	//response := grpcServer.RegisterService
}
