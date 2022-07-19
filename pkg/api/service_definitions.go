package api

type UserService interface {
	All() (users []User, err error)
	Get(uid int) (user User, err error)
	Delete(uid int) (err error)
	Update(user User) (err error)
}

type user_service struct {
	userRepo UserRepository
}

type BiddingService interface {
	All() (biddings []Bidding, err error)
	Get(bID int) (bid Bidding, err error)
	Delete(bID int) (err error)
	Update(bidding Bidding) (err error)
}

type bidding_service struct {
	biddingRepo BidRepository
}
