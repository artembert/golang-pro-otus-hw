package grpc

//go:generate protoc event_service.proto -I . --go_out=../../pkg/api/grpc/ --go-grpc_out=../../pkg/api/grpc/
