package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const port = "9000"

func main(){
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("failed to listen:", err)
	}

	grpcServer := grpc.NewServer()

	log.Println("start gRPC server on:", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve:", err)
	}
}
