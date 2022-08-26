package app

// bids

type CreateBidRequest struct {
	BidderID  int `json:"bidder_id"`
	BiddingID int `json:"bidding_id"`
	Amount    int `json:"amount"`
}
