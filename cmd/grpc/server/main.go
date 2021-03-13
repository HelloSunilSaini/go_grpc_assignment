package main

import (
	"log"
	"net"
	"os"

	"github.com/hellosunilsaini/go_grpc_assignment/controller"
	mygrpc "github.com/hellosunilsaini/go_grpc_assignment/grpc"
	"github.com/hellosunilsaini/go_grpc_assignment/repository/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// configure our user service
	repository := user.NewMockRepository()
	userService := user.NewService(repository)
	// configure our gRPC user service controller
	userServiceController := controller.NewUserServiceController(*userService)
	// start a gRPC server
	server := grpc.NewServer()
	mygrpc.RegisterUserServiceServer(server, userServiceController)

	reflection.Register(server)
	con, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		panic(err)
	}

	log.Printf("Server Started on Port : %v \n", os.Getenv("GRPC_PORT"))
	err = server.Serve(con)
	if err != nil {
		panic(err)
	}
}
