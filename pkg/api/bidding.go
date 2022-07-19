package api

func NewBiddingService(biddingRepo BidRepository) BiddingService {
	return &bidding_service{
		biddingRepo: biddingRepo,
	}
}

func (b *bidding_service) All() (biddings []Bidding, err error) {
	return
}

func (b *bidding_service) Get(bID int) (bidding Bidding, err error) {
	return
}

func (b *bidding_service) Delete(bID int) (err error) {
	return
}

func (b *bidding_service) Update(bidding Bidding) (err error) {
	return
}
