package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"

	"google.golang.org/grpc"

	"github.com/kubaceg/bookstore/internal/common/eventbus"
	"github.com/kubaceg/bookstore/internal/common/genproto/book"
	"github.com/kubaceg/bookstore/internal/common/genproto/reservation"
	"github.com/kubaceg/bookstore/internal/common/log"
	"github.com/kubaceg/bookstore/internal/common/server"
	"github.com/kubaceg/bookstore/internal/common/uuid"
	reservationsRepo "github.com/kubaceg/bookstore/internal/reservations/repository"
	"github.com/kubaceg/bookstore/internal/reservations/service"
)

var port, bookServiceAddress, rabbitMqUri, rabbitExchange string

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	port = os.Getenv("PORT")
	bookServiceAddress = os.Getenv("BOOK_CLIENT_URI")
	rabbitMqUri = os.Getenv("RABBITMQ_URI")
	rabbitExchange = os.Getenv("RABBITMQ_EXCHANGE")

	conn, err := grpc.Dial(bookServiceAddress, grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	bookClient := book.NewBookServiceClient(conn)
	logger := log.StdOut{}
	uuidGen := uuid.V4Generator{}
	repository := reservationsRepo.NewReservationInmemoryRepository()
	grpcServer := server.NewGrpcServer(logger)

	rabbitConn, err := amqp.Dial(rabbitMqUri)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer rabbitConn.Close()

	ch, err := rabbitConn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	rabbitPublisher := eventbus.NewRabbitMqPublisher(ch, rabbitExchange)

	grpcService := service.NewReservationsGrpcService(repository, bookClient, logger, &uuidGen, rabbitPublisher)

	grpcServer.RunGRPCServer(context.Background(), port, func(server *grpc.Server) {
		reservation.RegisterReservationServiceServer(server, grpcService)
	})
}

func failOnError(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
