package store

import (
	"reflect"
	"testing"
	"trm/internal/model"
)

func TestSearchUser(t *testing.T) {
	type args struct {
		userName string
	}
	tests := []struct {
		name string
		args args
		want *model.User
	}{
		{
			name: "existing user",
			args: args{
				userName: "user1",
			},
			want: &model.User{
				ID:       int64(1),
				Login:    "user1",
				Password: "pass1",
				UserType: UserTypeM,
			},
		},
		{
			name: "not existing user",
			args: args{
				userName: "user3",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindUser(tt.args.userName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
