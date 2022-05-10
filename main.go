package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	grpcserver "github.com/jha929/lego-grpc/internal/grpc_server"
	"github.com/jha929/lego-grpc/internal/repository"
	userpb "github.com/jha929/lego-grpc/protos/user"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const grpcPort = "9000"

func init() {
	godotenv.Load(".env")
}

func main() {
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_ADDR := os.Getenv("DB_ADDR")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")

	// Connect to DB (wf gorm)
	dsnStr := "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local"
	dbConnStr := fmt.Sprintf(
		dsnStr,
		DB_USERNAME,
		DB_PASSWORD,
		DB_ADDR,
		DB_PORT,
		DB_NAME,
	)
	db, err := gorm.Open(mysql.Open(dbConnStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&repository.User{})

	// Check gPRC Port
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatal("failed to listen on:", err)
	}

	// Create gRPC Server
	grpcServer := grpc.NewServer()
	userRepository, _ := repository.NewUserRepository(db)
	userpb.RegisterUserServer(
		grpcServer,
		&grpcserver.UserGrpcServer{UserRepository: userRepository},
	)

	// Start gRPC Server
	log.Println("start gRPC server on:", grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve:", err)
	}
}
