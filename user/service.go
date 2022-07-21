package user

import (
	"fmt"

	"go-mockery/user/model"
)

type Service struct {
	repo model.Repo
}

func NewService(repo model.Repo) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) InsertUser(user1 *model.User) error {

	err := s.repo.InsertUser(*user1)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) GetUser() ([]model.User, error) {

	userList, err := s.repo.GetUser()
	if err != nil {
		return nil, fmt.Errorf("failed to list users:  ")
	}

	return userList, nil
}
