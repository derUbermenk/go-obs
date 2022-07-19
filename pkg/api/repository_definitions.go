package api

type BidRepository interface {
	AllBiddings() (biddings []Bidding, err error)
	GetBidding(bID int) (bid Bidding, err error)
	DeleteBidding(bID int) (err error)
	UpdateBidding(bID int) (err error)
}
