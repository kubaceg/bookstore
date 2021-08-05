module github.com/kubaceg/bookstore/internal/reservations

go 1.15

replace github.com/kubaceg/bookstore/internal/common => ../common

require (
	github.com/joho/godotenv v1.3.0
	github.com/kubaceg/bookstore/internal/common v0.0.0-00010101000000-000000000000
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.5.1
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
)
