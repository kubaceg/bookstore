package server

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func RunGRPCServer(registerServer func(server *grpc.Server)) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	RunGRPCServerOnAddr(addr, registerServer)
}

func RunGRPCServerOnAddr(addr string, registerServer func(server *grpc.Server)) {
	grpcServer := grpc.NewServer()
	registerServer(grpcServer)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
