package app_test

import (
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

func (b *mockBiddingService) Delete(bID int) (err error) {
	return
}

func (mU *mockBiddingService) Update(biddingID int, bidding api.Bidding) (err error) {
	_, exists := mU.biddingRepo[biddingID]

	if !exists {
		return &api.ErrNonExistentResource{}
	}

	return nil
}
