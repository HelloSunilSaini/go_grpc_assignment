package clients

import (
	"context"
	"time"

	"github.com/hellosunilsaini/go_grpc_assignment/domain"
	mygrpc "github.com/hellosunilsaini/go_grpc_assignment/grpc"
	"google.golang.org/grpc"
)

var defaultRequestTimeout = time.Second * 10

type grpcService struct {
	grpcClient mygrpc.UserServiceClient
}

// NewGRPCService creates a new gRPC user service connection using the specified connection string.
func NewGRPCService(connString string) (*grpcService, error) {
	conn, err := grpc.Dial(connString, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &grpcService{grpcClient: mygrpc.NewUserServiceClient(conn)}, nil
}

func (s *grpcService) GetUsersByIDs(ids []int64) (result []domain.User) {
	result = []domain.User{}
	req := &mygrpc.GetUsersRequest{
		Ids: ids,
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	resp, err := s.grpcClient.GetUsers(ctx, req)
	if err != nil {
		return
	}
	for _, grpcUser := range resp.GetUsers() {
		u := unmarshalUser(grpcUser)
		result = append(result, *u)
	}
	return
}

func (s *grpcService) GetUserByID(id int64) (result *domain.User, err error) {
	req := &mygrpc.GetUserRequest{
		Id: id,
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	resp, err := s.grpcClient.GetUser(ctx, req)
	if err != nil {
		return
	}
	grpcUser := resp.GetUser()
	if grpcUser.GetId() == id {
		result = unmarshalUser(grpcUser)
	}
	return
}

func unmarshalUser(grpcUser *mygrpc.User) (result *domain.User) {
	result = &domain.User{}
	result.ID = grpcUser.Id
	result.FName = grpcUser.Fname
	result.City = grpcUser.City
	result.Height = grpcUser.Height
	result.Married = grpcUser.Married
	return
}
