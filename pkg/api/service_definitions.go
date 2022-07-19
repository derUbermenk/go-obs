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

type BiddingService interface {
	All() (biddings []Bidding, err error)
	Get(bID int) (bid Bidding, err error)
	Delete(bID int) (err error)
	Update(bID int) (err error)
}

type bidding_service struct {
	bidding_repo BidRepository
}
