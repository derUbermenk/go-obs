package app_test

import (
	"errors"
)

type mockBidService struct{}

func (mBid *mockBidService) CreateBid(bidderID, biddingID, bidAmount int) (bidID int, err error) {
	// bidder does not exist
	if bidderID > 1 {
		err = errors.New(`Bidder does not exist`)
		return
	}

	// bidding does not exist
	if biddingID > 1 {
		err = errors.New(`Bidding does not exist`)
		return
	}

	// invalid amount
	if bidAmount <= 100 {
		err = errors.New(`The bid amount is lower than lowest allowable bid`)
		return
	}

	// when all values pass return bidID 1
	bidID = 1
	return
}

func (mBid *mockBidService) UpdateBid(bidID, bidAmount int) (err error) {
	if bidAmount <= 100 {
		err = errors.New(`The bid amount is lower than lowest allowable bid`)
		return
	}

	if bidID != 1 {
		err = errors.New(`The bidder with the given bidding id does not exist`)
		return
	}

	return nil
}
