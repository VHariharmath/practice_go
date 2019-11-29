package main

import (
	"fmt"
	"log"
	"net"
	"practice_go/grpc/demo-grpc/api"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("Failed to open port: %v", err)
	}

	svc := &api.Server{}

	srv := grpc.NewServer()

	api.RegisterPingServer(srv, svc)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to server = %v", err)
	}
}
