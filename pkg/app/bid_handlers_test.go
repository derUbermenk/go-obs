package app_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"online-bidding-system/pkg/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SetUpBidHandlersTest() {
	bid_service = &mockBidService{}

	server = app.NewServer(
		router, user_service,
		bid_service,
		bidding_service,
		auth_service,
	)
}

func TearDownBidHandlersTest() {
	bid_service = nil
	server = nil
}

func TestCreateBid(t *testing.T) {
	SetUpRouter()
	SetUpBidHandlersTest()

	defer TearDownBidHandlersTest()
	defer TearDownRouter()

	var response *app.GenericResponse
	var expected_response *app.GenericResponse

	var request *http.Request
	var recorder *httptest.ResponseRecorder

	define_route(`POST`, `/bids`, server.CreateBid())

	// case 1: amount is lower than lowest allowable bid
	t.Run(
		`Amount is lower than lowest allowable bid`,
		func(t *testing.T) {

			createBidRequest := &app.CreateBidRequest{
				BidderID:  1,
				BiddingID: 1,
				Amount:    100,
			}
			jsonValue, _ := json.Marshal(createBidRequest)

			request = initialize_request(`POST`, `/bids`, bytes.NewBuffer(jsonValue))
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  false,
				Message: `The bid amount is lower than lowest allowable bid`,
			}

			json.Unmarshal(recorder.Body.Bytes(), &response)
			assert.Equal(t, http.StatusInternalServerError, recorder.Code)
			assert.Equal(t, response.Status, expected_response.Status)
			assert.Equal(t, response.Message, expected_response.Message)
			assert.Equal(t, response.Data, expected_response.Data)
		},
	)

	// case 2: bidder does not exist
	t.Run(
		`Bidder does not exist`,
		func(t *testing.T) {

			createBidRequest := &app.CreateBidRequest{
				BidderID:  2,
				BiddingID: 1,
				Amount:    100,
			}

			jsonValue, _ := json.Marshal(createBidRequest)

			request = initialize_request(`POST`, `/bids`, bytes.NewBuffer(jsonValue))
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  false,
				Message: `Bidder does not exist`,
			}

			json.Unmarshal(recorder.Body.Bytes(), &response)
			assert.Equal(t, http.StatusInternalServerError, recorder.Code)
			assert.Equal(t, response.Status, expected_response.Status)
			assert.Equal(t, response.Message, expected_response.Message)
			assert.Equal(t, response.Data, expected_response.Data)
		},
	)

	// case 3: bidding does not exist
	t.Run(
		`Bidding does not exist`,
		func(t *testing.T) {

			createBidRequest := &app.CreateBidRequest{
				BidderID:  1,
				BiddingID: 2,
				Amount:    100,
			}

			jsonValue, _ := json.Marshal(createBidRequest)

			request = initialize_request(`POST`, `/bids`, bytes.NewBuffer(jsonValue))
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  false,
				Message: `Bidding does not exist`,
			}

			json.Unmarshal(recorder.Body.Bytes(), &response)
			assert.Equal(t, http.StatusInternalServerError, recorder.Code)
			assert.Equal(t, response.Status, expected_response.Status)
			assert.Equal(t, response.Message, expected_response.Message)
			assert.Equal(t, response.Data, expected_response.Data)
		},
	)

	// case 4: All values are valid
	t.Run(
		`All values are valid`,
		func(t *testing.T) {
			createBidRequest := &app.CreateBidRequest{
				BidderID:  1,
				BiddingID: 1,
				Amount:    200,
			}

			jsonValue, _ := json.Marshal(createBidRequest)

			request = initialize_request(`POST`, `/bids`, bytes.NewBuffer(jsonValue))
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  true,
				Message: `Bid created`,
				Data:    1,
			}

			json.Unmarshal(recorder.Body.Bytes(), &response)
			assert.Equal(t, http.StatusCreated, recorder.Code)
			assert.Equal(t, response.Status, expected_response.Status)
			assert.Equal(t, response.Message, expected_response.Message)

			// from stack overflow answer
			// I am assuming: If you sent the JSON value through browser then
			// 		any number you sent that will be the type float64 so you cant get the value directly
			// 		int in golang.
			// 		https://stackoverflow.com/questions/18041334/convert-interface-to-int
			assert.Equal(t, int(response.Data.(float64)), expected_response.Data)
		},
	)
}

/*
func TestUpdateBid(t *testing.T) {
	SetUpRouter()
	SetUpBidHandlersTest()

	defer TearDownBidHandlersTest()
	defer TearDownRouter()

	var response *app.GenericResponse
	var expected_response *app.GenericResponse

	var request *http.Request
	var recorder *httptest.ResponseRecorder

	define_route(`PATCH`, `/bids/:bidID`, server.UpdateBid())

	// case 1: amount is lower than lowest allowable bid or current top bid amount
	t.Run(
		`Amount is lower than lowest allowable bid or current top bid amount`,
		func(t *testing.T) {
			updateBidRequest := &api.Bidding{
				Amount: 100,
			}

			jsonValue, _ := json.Marshal(updateBidRequest)

			request = initialize_request(`PATCH`, `/bids/1`, bytes.NewBuffer(jsonValue))
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  false,
				Message: `The bid amount is lower than lowest allowable bid`,
			}

			json.Unmarshal(recorder.Body.Bytes(), &response)
			assert.Equal(t, http.StatusBadRequest, recorder.Code)
			assert.Equal(t, response.Status, expected_response.Status)
			assert.Equal(t, response.Message, expected_response.Message)
			assert.Equal(t, response.Data, expected_response.Data)
		},
	)

	// case 2: all values pass
	t.Run(
		`All values pass`,
		func(t *testing.T) {
			updateBidRequest := &api.Bidding{
				Amount: 200,
			}

			jsonValue, _ := json.Marshal(updateBidRequest)

			request = initialize_request(`PATCH`, `/bids/1`, bytes.NewBuffer(jsonValue))
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  true,
				Message: `Bid updated`,
			}

			json.Unmarshal(recorder.Body.Bytes(), &response)
			assert.Equal(t, http.StatusBadRequest, recorder.Code)
			assert.Equal(t, response.Status, expected_response.Status)
			assert.Equal(t, response.Message, expected_response.Message)
			assert.Equal(t, response.Data, expected_response.Data)
		},
	)
}
*/
