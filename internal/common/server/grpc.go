package server

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/kubaceg/bookstore/internal/common/log"
)

type GrpcServer struct {
	log log.Logger
}

func NewGrpcServer(l log.Logger) *GrpcServer {
	return &GrpcServer{l}
}

func (g *GrpcServer) RunGRPCServer(ctx context.Context, port string, registerServer func(server *grpc.Server)) {
	grpcServer := grpc.NewServer()
	registerServer(grpcServer)

	addr := fmt.Sprintf(":%s", port)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		g.log.Fatal(ctx, err)
	}

	g.log.Info(ctx, fmt.Sprintf("Start service on %s port", port))

	if err := grpcServer.Serve(listen); err != nil {
		g.log.Fatal(ctx, err)
	}
}
