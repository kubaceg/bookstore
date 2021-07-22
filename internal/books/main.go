package main

import (
	"context"
	"flag"

	"google.golang.org/grpc"

	"github.com/kubaceg/bookstore/internal/books/repository"
	"github.com/kubaceg/bookstore/internal/books/service"
	"github.com/kubaceg/bookstore/internal/common/genproto/book"
	"github.com/kubaceg/bookstore/internal/common/log"
	"github.com/kubaceg/bookstore/internal/common/server"
)

var port string

func main() {
	flag.StringVar(&port, "port", "8080", "Service port")
	flag.Parse()

	logger := log.StdOut{}
	memoryRepo := repository.NewInMemoryBookRepository(logger)
	bookServer := service.NewBookGrpcService(memoryRepo)
	grpcServer := server.NewGrpcServer(logger)

	grpcServer.RunGRPCServer(context.Background(), port, func(server *grpc.Server) {
		book.RegisterBookServiceServer(server, bookServer)
	})
}
