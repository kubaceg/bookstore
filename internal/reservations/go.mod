module github.com/kubaceg/bookstore/internal/reservations

go 1.15

replace github.com/kubaceg/bookstore/internal/common => ../common

require (
	github.com/kubaceg/bookstore/internal/common v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.5.1
	golang.org/x/net v0.0.0-20201021035429-f5854403a974
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
)
