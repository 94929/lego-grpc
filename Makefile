prepare:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

gen:
	protoc -I=. \
	    --go_out . --go_opt paths=source_relative \
	    --go-grpc_out . --go-grpc_opt paths=source_relative \
	    protos/user/user.proto
run:
	go run main.go
