protobuf/pitcher.pb.go: protobuf/pitcher.proto
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative protobuf/pitcher.proto

proto: protobuf/pitcher.pb.go

all:
	proto
