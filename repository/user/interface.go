package user

import "github.com/hellosunilsaini/go_grpc_assignment/domain"

//Reader interface
type Reader interface {
	GetUserByID(userID int64) (*domain.User, error)
	GetUsersByIDs(userIDs []int64) map[int64]*domain.User
}

//Writer interface
type Writer interface {
}

//Filter interface
type Filter interface {
}

//Repository interface
type Repository interface {
	Reader
	Writer
	Filter
}
