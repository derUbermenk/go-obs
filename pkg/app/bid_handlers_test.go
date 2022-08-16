package app_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"online-bidding-system/pkg/api"
	"online-bidding-system/pkg/app"
	"testing"

	"github.com/go-playground/assert/v2"
)

func SetUpBidHandlersTest() {

}

func TearDownBidHandlersTest() {

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

			createBidRequest := &api.CreateBidRequest{
				ID:        1,
				BidderID:  1,
				BiddingId: 1,
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
			assert.Equal(t, http.StatusConflict, recorder.Code)
			assert.Equal(t, response.Status, expected_response.Status)
			assert.Equal(t, response.Message, expected_response.Message)
			assert.Equal(t, response.Data, expected_response.Data)
		},
	)

	// case 2: bidder does not exist
	t.Run(
		`Bidder does not exist`,
		func(t *testing.T) {

			createBidRequest := &api.CreateBidRequest{
				ID:        1,
				BidderID:  2,
				BiddingId: 1,
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
			assert.Equal(t, http.StatusNotFound, recorder.Code)
			assert.Equal(t, response.Status, expected_response.Status)
			assert.Equal(t, response.Message, expected_response.Message)
			assert.Equal(t, response.Data, expected_response.Data)
		},
	)

	// case 3: bidding does not exist
	t.Run(
		`Bidding does not exist`,
		func(t *testing.T) {

			createBidRequest := &api.CreateBidRequest{
				ID:        1,
				BidderID:  1,
				BiddingId: 2,
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
			assert.Equal(t, http.StatusNotFound, recorder.Code)
			assert.Equal(t, response.Status, expected_response.Status)
			assert.Equal(t, response.Message, expected_response.Message)
			assert.Equal(t, response.Data, expected_response.Data)
		},
	)

	// case 4: All values are valid
	t.Run(
		`All values are valid`,
		func(t *testing.T) {
			createBidRequest := &api.CreateBidRequest{
				ID:        1,
				BidderID:  1,
				BiddingId: 1,
				Amount:    100,
			}

			jsonValue, _ := json.Marshal(createBidRequest)

			request = initialize_request(`POST`, `/bids`, bytes.NewBuffer(jsonValue))
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  true,
				Message: `Bid Created`,
				Data:    1,
			}

			json.Unmarshal(recorder.Body.Bytes(), &response)
			assert.Equal(t, http.StatusOK, recorder.Code)
			assert.Equal(t, response.Status, expected_response.Status)
			assert.Equal(t, response.Message, expected_response.Message)
			assert.Equal(t, response.Data, expected_response.Data)
		},
	)
}

func TestUpdateBid(t *testing.T) {
	SetUpRouter()
	SetUpBidHandlersTest()

	defer TearDownBidHandlersTest()
	defer TearDownRouter()

	var response *app.GenericResponse
	var expected_response *app.GenericResponse

	var request *http.Request
	var recorder *httptest.ResponseRecorder

	define_route(`PUT`, `/bids/:bidID`, server.UpdateBid())

	// case 1: amount is lower than lowest allowable bid or current top bid amount
	t.Run(
		`Amount is lower than lowest allowable bid or current top bid amount`,
		func(t *testing.T) {
			updateBidRequest := &api.Bidding{
				Amount: 100,
			}

			jsonValue, _ := json.Marshal(updateBidRequest)

			request = initialize_request(`PUT`, `/bids/1`, bytes.NewBuffer(jsonValue))
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
}
