package api

type UserRepository interface {
	AllUsers() (users []User, err error)
	GetUser(uId int) (user []User, err error)
	DeleteUser(uID int) (err error)
	UpdateUser(user User) (err error)
}

type BidRepository interface {
	AllBiddings() (biddings []Bidding, err error)
	GetBidding(bID int) (bid Bidding, err error)
	DeleteBidding(bID int) (err error)
	UpdateBidding(bid Bidding) (err error)
}
