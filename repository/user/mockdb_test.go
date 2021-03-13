package user

import (
	"reflect"
	"testing"

	"github.com/hellosunilsaini/go_grpc_assignment/domain"
)

func TestMockRepository_GetUserByID(t *testing.T) {
	type args struct {
		userID int64
	}
	tests := []struct {
		name    string
		m       *MockRepository
		args    args
		want    *domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockRepository{}
			got, err := m.GetUserByID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("MockRepository.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockRepository.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockRepository_GetUsersByIDs(t *testing.T) {
	type args struct {
		userIDs []int64
	}
	tests := []struct {
		name string
		m    *MockRepository
		args args
		want []*domain.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockRepository{}
			if got := m.GetUsersByIDs(tt.args.userIDs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockRepository.GetUsersByIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}
