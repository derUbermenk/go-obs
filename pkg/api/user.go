package api

func NewUserService(userRepo UserRepository) UserService {
	return &user_service{
		userRepo: userRepo,
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

func (u *user_service) Update(user User) (err error) {
	return
}
