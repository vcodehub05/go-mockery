package model

type Service interface {
	InsertUser(user *User) error
	GetUser() ([]User, error)
}

type Repo interface {
	InsertUser(user User) error
	GetUser() ([]User, error)
}
