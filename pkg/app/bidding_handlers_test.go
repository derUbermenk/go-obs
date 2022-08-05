package app_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"online-bidding-system/pkg/api"
	"online-bidding-system/pkg/app"
	"testing"

	"github.com/go-playground/assert/v2"
)

func SetUpBiddingHandlersTest() {
	// initialize api variables
	bidding_service = &mockBiddingService{
		biddingRepo: biddingRepo,
	}

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

func TestAllBiddings(t *testing.T) {
	SetUpRouter()
	SetUpBiddingHandlersTest()

	defer TearDownBiddingHandlersTests()
	defer TearDownRouter()

	var response *app.GenericResponse
	var expected_response *app.GenericResponse

	var request *http.Request
	var recorder *httptest.ResponseRecorder

	define_route(`GET`, `/biddings`, server.AllBiddings())
	request = initialize_request(`GET`, `/biddings`, nil)
	recorder = send_request(request)

	// marshall and unmarshall the expected response to be able to check
	// for equality of the data returned.
	expected_byte_response, err := json.Marshal(&app.GenericResponse{
		Status:  true,
		Message: `Biddings successfully retrieved`,
		Data:    []api.Bidding{biddingRepo[0], biddingRepo[1]},
	})

	assert.Equal(t, err, nil)

	json.Unmarshal(expected_byte_response, &expected_response)
	json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expected_response.Status, response.Status)
	assert.Equal(t, expected_response.Message, response.Message)
	assert.Equal(t, expected_response.Data, response.Data)
}
