package app_test

import "online-bidding-system/pkg/app"

func SetUpBiddingHandlersTest() {
	// initialize api variables
	bidding_service = &mockBiddingService

	server = app.NewServer(
		router, user_service,
		bidding_service, auth_service,
	)
}

// sets the bidding_service and server to nil,
// technically resetting the values.
func TearDownBiddingHandlersTests() {
	bidding_service = nil
	server = nil
}
