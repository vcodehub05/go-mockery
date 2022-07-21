package repository

import (
	"fmt"

	"go-mockery/user/model"
)

type Repo struct {
	datamp map[int]model.User
}

func NewRepo() model.Repo {
	return Repo{datamp: make(map[int]model.User)}
}

func (r Repo) InsertUser(user model.User) error {
	r.datamp[int(user.ID)] = user
	fmt.Println(r.datamp)

	return nil

}

func (r Repo) GetUser() ([]model.User, error) {
	var users []model.User
	for key := range r.datamp {
		user := r.datamp[key]
		users = append(users, user)
	}

	if len(users) == 0 {
		err := fmt.Errorf("Empty")
		return nil, err
	}

	return users, nil
}
