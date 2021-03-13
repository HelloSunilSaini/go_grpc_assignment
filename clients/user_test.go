package clients

import (
	"reflect"
	"testing"

	"github.com/hellosunilsaini/go_grpc_assignment/domain"
	mygrpc "github.com/hellosunilsaini/go_grpc_assignment/grpc"
)

func Test_grpcService_GetUsersByIDs(t *testing.T) {
	type fields struct {
		grpcClient mygrpc.UserServiceClient
	}
	type args struct {
		ids []int64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult map[int64]*domain.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &grpcService{
				grpcClient: tt.fields.grpcClient,
			}
			if gotResult := s.GetUsersByIDs(tt.args.ids); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("grpcService.GetUsersByIDs() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_grpcService_GetUserByID(t *testing.T) {
	type fields struct {
		grpcClient mygrpc.UserServiceClient
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult *domain.User
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &grpcService{
				grpcClient: tt.fields.grpcClient,
			}
			gotResult, err := s.GetUserByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("grpcService.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("grpcService.GetUserByID() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
