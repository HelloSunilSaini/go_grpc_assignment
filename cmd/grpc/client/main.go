package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hellosunilsaini/go_grpc_assignment/clients"
)

func main() {
	service, err := clients.NewGRPCService(os.Getenv("GRPC_CONNECTION"))
	if err != nil {
		panic(err)
	}
	userIdsStr := os.Getenv("USER_IDS")
	userIdsListStr := strings.Split(userIdsStr, ",")
	userIdsList := []int64{}
	for _, userId := range userIdsListStr {
		u, err := strconv.Atoi(userId)
		if err == nil {
			userIdsList = append(userIdsList, int64(u))
		}
	}
	for _, i := range userIdsList {
		user, err := service.GetUserByID(int64(i))
		fmt.Printf("userId - %v, user - %v, err - %v\n", i, user, err)
	}
	users := service.GetUsersByIDs(userIdsList)
	fmt.Printf("userList - %v\n", users)
}
