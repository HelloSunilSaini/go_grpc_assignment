package user

import "github.com/hellosunilsaini/go_grpc_assignment/domain"

type Service struct {
	repo Repository
}

//NewService create new repository
func NewService(repository Repository) *Service {
	return &Service{
		repo: repository,
	}
}

func (s *Service) GetUserByID(userID int64) (*domain.User, error) {
	return s.repo.GetUserByID(userID)
}

func (s *Service) GetUsersByIDs(userIDs []int64) map[int64]*domain.User {
	return s.repo.GetUsersByIDs(userIDs)
}
