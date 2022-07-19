package api

import "online-bidding-system/pkg/repository"

func NewUserService(storage repository.Storage) UserService {
	return &user_service{
		user_repo: storage,
	}
}

func (u *user_service) All() (users []User, err error) {
	return
}

func (u *user_service) Get(uid int) (user User, err error) {
	return
}

func (u *user_service) Delete(uid int) (err error) {
	return
}

func (u *user_service) Update(User) (user User, err error) {
	return
}
