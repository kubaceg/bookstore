protoMessagePath = internal/common/genproto/message
protoServicePath = internal/common/genproto/service

.PHONY: proto
proto:
	protoc -I /usr/local/include -I api/protobuf --go_out=$(protoMessagePath) --go-grpc_out=$(protoServicePath) api/protobuf/book/book.proto
	protoc -I /usr/local/include -I api/protobuf --go_out=$(protoMessagePath) --go-grpc_out=$(protoServicePath) api/protobuf/user/user.proto
	protoc -I /usr/local/include -I api/protobuf --go_out=$(protoMessagePath) --go-grpc_out=$(protoServicePath) api/protobuf/reservation/reservation.proto




