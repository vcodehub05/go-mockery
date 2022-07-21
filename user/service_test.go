package user_test

import (
	"errors"
	"testing"

	"go-mockery/user"
	"go-mockery/user/model"
	"go-mockery/user/model/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var err = errors.New("error in service")

func TestService_Insert(t *testing.T) {
	t.Parallel()
	type fields struct {
		repo model.Repo
	}
	type args struct {
		user *model.User
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "insert sucessful",
			fields: fields{
				repo: func() model.Repo {
					repo := new(mocks.Repo)
					repo.On("InsertUser", mock.Anything).Return(nil)
					return repo
				}(),
			},
			args: args{
				user: &model.User{
					ID:        1,
					FirstName: "swati",
					LastName:  "singh",
				},
			},
			wantErr: false,
		},
		{
			name: "insert unsucessful",
			fields: fields{
				repo: func() model.Repo {
					repo := new(mocks.Repo)
					repo.On("InsertUser", mock.Anything).Return(err)
					return repo
				}(),
			},
			args: args{
				user: &model.User{
					ID:        1,
					FirstName: "swati",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := user.NewService(tt.fields.repo)
			if err := s.InsertUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("InsertUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_GetUser(t *testing.T) {
	t.Parallel()
	type fields struct {
		repo model.Repo
	}
	tests := []struct {
		name    string
		fields  fields
		want    []model.User
		wantErr bool
	}{
		{
			name: "get sucessful",
			fields: fields{
				repo: func() model.Repo {
					repo := new(mocks.Repo)
					repo.On("GetUser").Return([]model.User{
						{
							ID:        1,
							FirstName: "Jhon",
							LastName:  "caloun",
						},
						{
							ID:        2,
							FirstName: "sam",
							LastName:  "rio",
						},
					}, nil)
					return repo
				}(),
			},
			wantErr: false,
			want: []model.User{
				{
					ID:        1,
					FirstName: "Jhon",
					LastName:  "caloun",
				},
				{
					ID:        2,
					FirstName: "sam",
					LastName:  "rio",
				},
			},
		},
		{
			name: "get unsucessful",
			fields: fields{
				repo: func() model.Repo {
					repo := new(mocks.Repo)
					repo.On("GetUser").Return(nil, err)
					return repo
				}(),
			},
			wantErr: true,
			want:    nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := user.NewService(tt.fields.repo)
			if c, err := s.GetUser(); (err != nil) != tt.wantErr || !assert.Equal(t, c, tt.want) {
				t.Errorf("GetUser() error = %v, wantErr %v, user %v, want %v", err, tt.wantErr, c, tt.want)
			}
		})
	}
}
