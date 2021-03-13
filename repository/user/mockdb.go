package user

import (
	"github.com/hellosunilsaini/go_grpc_assignment/domain"
)

type MockRepository struct{}

func NewMockRepository() *MockRepository {
	return &MockRepository{}
}

var MockData = map[int64]domain.User{
	1: {
		ID:      1,
		FName:   "Steve",
		Phone:   "1234567890",
		Height:  5.8,
		Married: true,
	},
	2: {
		ID:      2,
		FName:   "Sunil",
		Phone:   "1231231231",
		Height:  5.8,
		Married: false,
	},
	3: {
		ID:      3,
		FName:   "Jhon",
		Phone:   "2342342341",
		Height:  5.3,
		Married: false,
	},
	4: {
		ID:      4,
		FName:   "Mike",
		Phone:   "1234123412",
		Height:  5.7,
		Married: true,
	},
	5: {
		ID:      5,
		FName:   "Haward",
		Phone:   "1234512345",
		Height:  5.6,
		Married: true,
	},
	6: {
		ID:      6,
		FName:   "Bob",
		Phone:   "1223341223",
		Height:  5.5,
		Married: false,
	},
	7: {
		ID:      7,
		FName:   "Alice",
		Phone:   "1233211233",
		Height:  5.3,
		Married: false,
	},
}

// GetUserByID function returns single user object by userId if found
// else returns error
func (m *MockRepository) GetUserByID(userID int64) (*domain.User, error) {
	user, ok := MockData[userID]
	if ok {
		return &user, nil
	}
	return nil, domain.NotFoundError
}

// GetUsersByIDs function returns slice of user objects by slice of userIds
// it will return empty slice if no userId found
func (m *MockRepository) GetUsersByIDs(userIDs []int64) map[int64]*domain.User {
	userMap := map[int64]*domain.User{}
	for _, userID := range userIDs {
		user, ok := MockData[userID]
		if ok {
			userMap[userID] = &user
		}
	}
	return userMap
}
