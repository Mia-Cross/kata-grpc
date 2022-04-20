package main

import (
	"google.golang.org/grpc"
	pizza "kata_grpc/api/proto"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	service := pizza.PizzaServiceServer()
	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	pizza.RegisterPizzaServiceServer(grpcServer, service)
}
