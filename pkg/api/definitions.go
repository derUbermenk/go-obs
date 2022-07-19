package api

import "online-bidding-system/pkg/repository"

type UserService interface {
	All() (users []User, err error)
	Get(uid int) (user User, err error)
	Delete(uid int) (err error)
	Update(User) (user User, err error)
}

type user_service struct {
	user_repo repository.Storage
}

type User struct{}
