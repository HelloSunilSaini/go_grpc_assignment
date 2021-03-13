package main

import (
	"fmt"
	"os"

	"github.com/hellosunilsaini/go_grpc_assignment/clients"
)

func main() {
	service, err := clients.NewGRPCService(os.Getenv("GRPC_CONNECTION"))
	if err != nil {
		panic(err)
	}
	for i := 1; i <= 10; i++ {
		user, err := service.GetUserByID(int64(i))
		fmt.Printf("userId - %v, user - %v, err - %v\n", i, user, err)
	}
	users := service.GetUsersByIDs([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Printf("userList - %v\n", users)
}
