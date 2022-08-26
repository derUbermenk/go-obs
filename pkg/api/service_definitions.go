package api

type UserService interface {
	All() (users []User, err error)
	Get(id int) (user User, err error)
	Delete(id int) (err error)
	Update(id int, user User) (err error)
}

type user_service struct {
	userRepo UserRepository
}

type BiddingService interface {
	All() (biddings []Bidding, err error)
	Get(bID int) (bid Bidding, err error)
	Delete(bID int) (err error)
	Update(id int, bidding Bidding) (err error)
}

type bidding_service struct {
	biddingRepo BidRepository
}

type BidService interface {
	CreateBid(bidderID, biddingID, bidAmount int) (bidID int, err error)
	UpdateBid(bidID, bidAmount int) (err error)
}

type bid_service struct{}

type AuthService interface {
	ValidateCredentials(email, password string) (validity bool, err error)
	GenerateAccessToken(email string, expiration int64) (signed_access_token string, err error)
	GenerateRefreshToken(email string, customKey string) (signed_refresh_token string, err error)
	ValidateAccessToken(access_token string) (status int)
	ValidateRefreshToken(refresh_token, custom_key string) (validity bool)
}

type authentication_service struct {
	authRepo AuthRepository
}
