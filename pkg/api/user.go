package api

func NewUserService(userRepo UserRepository) UserService {
	return &user_service{
		userRepo: userRepo,
	}
}

func (u *user_service) All() (users []User, err error) {
	return
}

func (u *user_service) Get(id int) (user User, err error) {
	return
}

func (u *user_service) Delete(id int) (err error) {
	return
}

func (u *user_service) Update(id int, user User) (err error) {
	return
}

func (u *user_service) GetBids(id string) (bids []Bid, err error) {
	return
}
