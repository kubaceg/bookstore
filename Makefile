.PHONY: proto
proto:
	protoc -I /usr/local/include -I api/protobuf --go_out=internal/common/genproto api/protobuf/book/book.proto
	protoc -I /usr/local/include -I api/protobuf --go_out=internal/common/genproto api/protobuf/user/user.proto
	protoc -I /usr/local/include -I api/protobuf --go_out=internal/common/genproto api/protobuf/reservation/reservation.proto




