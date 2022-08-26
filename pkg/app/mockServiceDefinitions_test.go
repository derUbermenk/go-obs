package app_test

import (
	"errors"
	"online-bidding-system/pkg/api"
)

// this must be the same as that defined in api definition of the user.
// this is made so that no api code is called in this test
type mockUserService struct {
	userRepo map[int]api.User
}

var userRepo = map[int]api.User{
	0: {
		Name:           "User One",
		Email:          "user1@email.com",
		HashedPassword: "x120asd",
	},
	1: {
		Name:           "User Two",
		Email:          "user2@email.com",
		HashedPassword: "y562ash",
	},
}

func (mU *mockUserService) All() ([]api.User, error) {
	// return an array of all the users
	return []api.User{mU.userRepo[0], mU.userRepo[1]}, nil
}

func (mU *mockUserService) Get(userID int) (api.User, error) {
	user, exists := mU.userRepo[userID]

	if !exists {
		return user, &api.ErrNonExistentResource{}
	}

	return user, nil
}

func (mU *mockUserService) Delete(userID int) error {
	// simulate a successful delete operation
	_, exists := mU.userRepo[userID]

	if !exists {
		return &api.ErrNonExistentResource{}
	}

	return nil
}

func (mU *mockUserService) Update(userID int, user api.User) error {
	_, exists := mU.userRepo[userID]

	if !exists {
		return &api.ErrNonExistentResource{}
	}

	return nil
}

type mockBiddingService struct {
	biddingRepo map[int]api.Bidding
}

var biddingRepo = map[int]api.Bidding{
	0: {},
	1: {},
}

func (mB *mockBiddingService) All() (biddings []api.Bidding, err error) {
	return []api.Bidding{
		biddingRepo[0],
		biddingRepo[1],
	}, nil
}

func (mB *mockBiddingService) Get(biddingID int) (bidding api.Bidding, err error) {
	bidding, exists := mB.biddingRepo[biddingID]

	if !exists {
		return bidding, &api.ErrNonExistentResource{}
	}

	return bidding, nil
}

func (mB *mockBiddingService) Delete(biddingID int) (err error) {
	// simulate a successful delete operation
	_, exists := mB.biddingRepo[biddingID]

	if !exists {
		return &api.ErrNonExistentResource{}
	}

	return nil
}

func (mB *mockBiddingService) Update(biddingID int, bidding api.Bidding) (err error) {
	_, exists := mB.biddingRepo[biddingID]

	if !exists {
		return &api.ErrNonExistentResource{}
	}

	return nil
}

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

func (mBid *mockBidService) UpdateBid(bidID, bidAmount int) (updatedBidID int, err error) {
	return
}
