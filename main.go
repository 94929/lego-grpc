package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

    userpb "github.com/jha929/lego-grpc/protos/user"
    "github.com/jha929/lego-grpc/internal/server"
)

const port = "9000"

func main(){
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
    userpb.RegisterUserServer(grpcServer, &server.UserPbServer{})

	log.Println("start gRPC server on:", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve:", err)
	}
}
