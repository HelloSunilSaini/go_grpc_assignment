package controller

import (
	"context"
	"reflect"
	"testing"

	mygrpc "github.com/hellosunilsaini/go_grpc_assignment/grpc"
	"github.com/hellosunilsaini/go_grpc_assignment/repository/user"
)

func Test_userServiceController_GetUsers(t *testing.T) {
	type fields struct {
		userService user.Service
	}
	type args struct {
		ctx context.Context
		req *mygrpc.GetUsersRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *mygrpc.GetUsersResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctlr := &userServiceController{
				userService: tt.fields.userService,
			}
			got, err := ctlr.GetUsers(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("userServiceController.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userServiceController.GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userServiceController_GetUser(t *testing.T) {
	type fields struct {
		userService user.Service
	}
	type args struct {
		ctx context.Context
		req *mygrpc.GetUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *mygrpc.GetUserResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctlr := &userServiceController{
				userService: tt.fields.userService,
			}
			got, err := ctlr.GetUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("userServiceController.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userServiceController.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
