package app_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"online-bidding-system/pkg/api"
	"online-bidding-system/pkg/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

// this must be the same as that defined in api definition of the user.
// this is made so that no api code is called in this test
type mockBiddingService struct {
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

func SetUpBiddingHandlersTest() {
	// initialize api variables
	bidding_service = &mockBiddingService{
		biddingRepo: biddingRepo,
	}

	server = app.NewServer(
		router, user_service,
		bid_service, bidding_service,
		auth_service,
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

func TestGetBidding(t *testing.T) {
	SetUpRouter()
	SetUpBiddingHandlersTest()

	defer TearDownBiddingHandlersTests()
	defer TearDownRouter()

	var response *app.GenericResponse
	var expected_response *app.GenericResponse

	var request *http.Request
	var recorder *httptest.ResponseRecorder

	define_route(`GET`, `/biddings/:id`, server.GetBidding())

	t.Run(
		"Bidding does not exist",
		func(t *testing.T) {
			request = initialize_request(`GET`, `/biddings/2`, nil)
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  false,
				Message: "Bidding does not exist",
				Data:    nil,
			}

			json.Unmarshal(recorder.Body.Bytes(), &response)

			assert.Equal(t, http.StatusNotFound, recorder.Code)
			assert.Equal(t, response.Status, expected_response.Status)
			assert.Equal(t, response.Message, expected_response.Message)
			assert.Equal(t, response.Data, expected_response.Data)
		},
	)

	t.Run(
		"Bidding Exists",
		func(t *testing.T) {
			request = initialize_request(`GET`, `/biddings/1`, nil)
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  true,
				Message: "Bidding retrieved",
				Data:    biddingRepo[1],
			}

			expected_json_response, err := json.Marshal(expected_response)
			assert.Equal(t, err, nil)

			json.Unmarshal(recorder.Body.Bytes(), &response)
			json.Unmarshal(expected_json_response, &expected_response)

			assert.Equal(t, http.StatusFound, recorder.Code)
			assert.Equal(t, response.Status, expected_response.Status)
			assert.Equal(t, response.Message, expected_response.Message)
			assert.Equal(t, response.Data, expected_response.Data)
		},
	)
}

func TestUpdateBidding(t *testing.T) {
	SetUpRouter()
	SetUpBiddingHandlersTest()

	defer TearDownBiddingHandlersTests()
	defer TearDownRouter()

	var response *app.GenericResponse
	var expected_response *app.GenericResponse

	define_route(`PATCH`, `/biddings/:id/update`, server.UpdateBidding())

	var request *http.Request
	var recorder *httptest.ResponseRecorder

	// the updated fields are not necessary
	// we only care about the response
	bidding := api.Bidding{}

	// case 1. user does not exist
	t.Run(
		"Bidding does not exist",
		func(t *testing.T) {
			jsonValue, _ := json.Marshal(bidding)
			request = initialize_request(`PATCH`, `/biddings/2/update`, bytes.NewBuffer(jsonValue))
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  false,
				Message: `Bidding does not exist`,
			}

			json.Unmarshal(recorder.Body.Bytes(), &response)

			assert.Equal(t, http.StatusBadRequest, recorder.Code)
			assert.Equal(t, expected_response.Data, response.Data)
			assert.Equal(t, expected_response.Message, response.Message)
		},
	)

	t.Run(
		"Bidding exists",
		func(t *testing.T) {
			jsonValue, _ := json.Marshal(bidding)
			request = initialize_request(`PATCH`, `/biddings/1/update`, bytes.NewBuffer(jsonValue))
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  true,
				Message: `Bidding successfully updated`,
			}

			json.Unmarshal(recorder.Body.Bytes(), &response)

			assert.Equal(t, http.StatusOK, recorder.Code)
			assert.Equal(t, expected_response.Data, response.Data)
			assert.Equal(t, expected_response.Message, response.Message)
		},
	)
}

func TestDeleteBidding(t *testing.T) {
	SetUpRouter()
	SetUpBiddingHandlersTest()

	defer TearDownBiddingHandlersTests()
	defer TearDownRouter()

	var response *app.GenericResponse
	var expected_response *app.GenericResponse

	var request *http.Request
	var recorder *httptest.ResponseRecorder

	define_route(`DELETE`, `/biddings/:id/delete`, server.DeleteBidding())

	t.Run(
		"Bidding does not exist",
		func(*testing.T) {
			request = initialize_request(`DELETE`, `/biddings/2/delete`, nil)
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  false,
				Message: `Bidding does not exist`,
			}

			json.Unmarshal(recorder.Body.Bytes(), &response)

			assert.Equal(t, http.StatusBadRequest, recorder.Code)
			assert.Equal(t, expected_response.Status, response.Status)
			assert.Equal(t, expected_response.Message, response.Message)
		},
	)

	t.Run(
		"Bidding exists",
		func(*testing.T) {
			request = initialize_request(`DELETE`, `/biddings/1/delete`, nil)
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  true,
				Message: `Bidding successfully deleted`,
			}

			json.Unmarshal(recorder.Body.Bytes(), &response)

			assert.Equal(t, http.StatusOK, recorder.Code)
			assert.Equal(t, expected_response.Status, response.Status)
			assert.Equal(t, expected_response.Message, response.Message)
		},
	)
}
