package main

import (
	"net"
	"os"

	"github.com/hellosunilsaini/go_grpc_assignment/controller"
	mygrpc "github.com/hellosunilsaini/go_grpc_assignment/grpc"
	"github.com/hellosunilsaini/go_grpc_assignment/repository/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// configure our core service
	repository := user.NewMockRepository()
	userService := user.NewService(repository)
	// configure our gRPC service controller
	userServiceController := controller.NewUserServiceController(*userService)
	// start a gRPC server
	server := grpc.NewServer()
	mygrpc.RegisterUserServiceServer(server, userServiceController)

	reflection.Register(server)
	con, err := net.Listen("tcp", os.Getenv("GRPC_ADDR"))
	if err != nil {
		panic(err)
	}

	err = server.Serve(con)
	if err != nil {
		panic(err)
	}
}
