package main

import (
	"log"
	"net"

	"github.com/HarrisonHenri/full-cycle-grpc/pb"
	"github.com/HarrisonHenri/full-cycle-grpc/services"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	// reflection.Register(grpcServer)

	if err != grpcServer.Serve(lis) {
		log.Fatalf("Could not serve: %v", err)
	}
}
