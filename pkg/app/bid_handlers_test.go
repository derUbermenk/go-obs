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
}
