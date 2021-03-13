package controller

import (
	"context"

	"github.com/hellosunilsaini/go_grpc_assignment/domain"
	mygrpc "github.com/hellosunilsaini/go_grpc_assignment/grpc"
	"github.com/hellosunilsaini/go_grpc_assignment/repository/user"
)

// userServiceController implements the gRPC UserServiceServer interface.
type userServiceController struct {
	userService user.Service
}

// NewUserServiceController instantiates a new UserServiceServer.
func NewUserServiceController(userService user.Service) mygrpc.UserServiceServer {
	return &userServiceController{
		userService: userService,
	}
}

// GetUsers calls the core service's GetUsersByIDs method and maps the result to a grpc service response.
func (ctlr *userServiceController) GetUsers(ctx context.Context, req *mygrpc.GetUsersRequest) (*mygrpc.GetUsersResponse, error) {
	resultMap := ctlr.userService.GetUsersByIDs(req.GetIds())

	resp := &mygrpc.GetUsersResponse{}
	for _, u := range resultMap {
		resp.Users = append(resp.Users, marshalUser(u))
	}
	return resp, nil
}

// GetUser calls the core service's GetUserByID method and maps the result to a grpc service response.
func (ctlr *userServiceController) GetUser(ctx context.Context, req *mygrpc.GetUserRequest) (*mygrpc.GetUserResponse, error) {
	user, err := ctlr.userService.GetUserByID(req.GetId())
	if err != nil {
		return nil, err
	}
	resp := &mygrpc.GetUserResponse{}
	resp.User = marshalUser(user)
	return resp, nil
}

// marshalUser marshals a domain User object into a gRPC layer User.
func marshalUser(u *domain.User) *mygrpc.User {
	return &mygrpc.User{
		Id:      u.ID,
		Fname:   u.FName,
		City:    u.City,
		Phone:   u.Phone,
		Height:  u.Height,
		Married: u.Married,
	}
}
