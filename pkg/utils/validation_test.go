package utils

import (
	"reflect"
	"testing"

	"github.com/supperdoggy/helper/pkg/models"
	"github.com/supperdoggy/helper/pkg/models/dbmodels"
)

func TestValidateUserEmailAndPassword(t *testing.T) {
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "not ok email",
			args: args{
				email:    "asjdklsajd",
				password: "Ssdz@1sdv",
			},
			wantErr: true,
		},
		{
			name: "not okay password",
			args: args{
				email:    "email@mail.com",
				password: "1",
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				email:    "mail@mail.com",
				password: "Ssdz@1sdv",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateUserEmailAndPassword(tt.args.email, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("ValidateUserEmailAndPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMapDBUserToResponseUser(t *testing.T) {
	type args struct {
		u dbmodels.User
	}
	tests := []struct {
		name string
		args args
		want *models.User
	}{
		{
			name: "ok",
			args: args{
				u: dbmodels.User{
					ID:        "someid",
					Email:     "email@mail.com",
					CreatedAt: 123213,
					EditedAt:  213213,
				},
			},
			want: &models.User{
				ID:        "someid",
				Email:     "email@mail.com",
				CreatedAt: 123213,
				EditedAt:  213213,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapDBUserToResponseUser(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapDBUserToResponseUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
