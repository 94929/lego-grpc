grpcgen:
	protoc -I=. \
	    --go_out . --go_opt paths=source_relative \
	    --go-grpc_out . --go-grpc_opt paths=source_relative \
	    protos/user/user.proto

gqlgen:
	go env

run:
	go run main.go
