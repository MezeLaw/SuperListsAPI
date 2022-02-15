package service

import (
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestNewUserListService(t *testing.T) {
	type args struct {
		repository IUserListRepository
	}
	tests := []struct {
		name string
		args args
		want UserListService
	}{
		{
			name: "Service with nil repo should pass",
			args: args{repository: nil},
			want: NewUserListService(nil),
		},
		{
			name: "Service with no nil repo should pass",
			args: args{repository: NewMockIUserListRepository(gomock.NewController(t))},
			want: NewUserListService(NewMockIUserListRepository(gomock.NewController(t))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserListService(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserListService() = %v, want %v", got, tt.want)
			}
		})
	}
}
