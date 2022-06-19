package main

import (
	"authorization/authorization"
	server "authorization/grpcserver"
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		panic("failed to start application, could not load environment configuration.")
	}
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Info("authorizaion server started listening")
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)
	authorization.RegisterAuthorizationServer(grpcServer, &server.Server{})
	grpcServer.Serve(lis)
}
