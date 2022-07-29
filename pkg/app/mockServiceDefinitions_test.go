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
	return mU.userRepo[userID], nil
}

func (mU *mockUserService) Delete(userID int) error {
	// simulate a successful delete operation
	_, exists := mU.userRepo[userID]

	if !exists {
		return &api.ErrNonExistentUser{}
	}

	return nil
}

func (mU *mockUserService) Update(userID int, user api.User) error {
	_, exists := mU.userRepo[userID]

	if !exists {
		return &api.ErrNonExistentUser{}
	}

	return nil
}
