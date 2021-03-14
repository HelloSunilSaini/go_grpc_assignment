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
		{
			name: "test 1",
			m:    NewMockRepository(),
			args: args{userID: 1},
			want: &domain.User{
				ID:      1,
				FName:   "Steve",
				Phone:   "1234567890",
				Height:  5.8,
				Married: true,
			},
			wantErr: false,
		},
		{
			name:    "test 2",
			m:       NewMockRepository(),
			args:    args{userID: 10},
			want:    nil,
			wantErr: true,
		},
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
		{
			name: "test 1",
			m:    NewMockRepository(),
			args: args{userIDs: []int64{1, 2, 3, 30}},
			want: []*domain.User{
				{
					ID:      1,
					FName:   "Steve",
					Phone:   "1234567890",
					Height:  5.8,
					Married: true,
				},
				{
					ID:      2,
					FName:   "Sunil",
					Phone:   "1231231231",
					Height:  5.8,
					Married: false,
				},
				{
					ID:      3,
					FName:   "Jhon",
					Phone:   "2342342341",
					Height:  5.3,
					Married: false,
				},
			},
		},
		{
			name: "test 2",
			m:    NewMockRepository(),
			args: args{userIDs: []int64{11, 12, 13}},
			want: []*domain.User{},
		},
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
