package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	db "tcpserver/database"
	"tcpserver/proto"
	"tcpserver/repository"
	"tcpserver/service"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	DB := db.InitDB()
	grpcServer := grpc.NewServer()
	proto.RegisterMyAppServer(grpcServer, &service.Server{UserRepo: &repository.UserRepositoryImpl{DB: DB}})
	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("server listening on 50051")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
