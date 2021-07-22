GENPROTO_PATH = internal/common/
MODULE_NAME=github.com/kubaceg/bookstore/internal/common
GO_OPT_FLAG=--go_opt=module=${MODULE_NAME}
GRPC_OPT_FLAG=--go-grpc_opt=module=${MODULE_NAME}

.PHONY: proto
proto:
	protoc -I /usr/local/include -I api/protobuf $(GO_OPT_FLAG) $(GRPC_OPT_FLAG) --go_out=$(GENPROTO_PATH) --go-grpc_out=$(GENPROTO_PATH) api/protobuf/book/book.proto
	protoc -I /usr/local/include -I api/protobuf $(GO_OPT_FLAG) $(GRPC_OPT_FLAG) --go_out=$(GENPROTO_PATH) --go-grpc_out=$(GENPROTO_PATH) api/protobuf/user/user.proto
	protoc -I /usr/local/include -I api/protobuf $(GO_OPT_FLAG) $(GRPC_OPT_FLAG) --go_out=$(GENPROTO_PATH) --go-grpc_out=$(GENPROTO_PATH) api/protobuf/reservation/reservation.proto

.PHONY: tests
tests:
	cd internal/books && go test -count 5 ./...
