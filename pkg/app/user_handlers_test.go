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

// sets up server to use mock user service
// for all user handler tests
func SetUpUserHandlersTest() {
	// initialize api variables
	user_service = &mockUserService{userRepo: userRepo}

	// initialize the server using the initialized services
	server = app.NewServer(
		router, user_service,
		bidding_service, auth_service,
	)
}

// sets the user_service and server to nil,
// essentially resetting the values.
func TearDownUserHandlersTests() {
	user_service = nil
	server = nil
}

func TestAllUsers(t *testing.T) {
	SetUpRouter()
	SetUpUserHandlersTest()

	defer TearDownUserHandlersTests()
	defer TearDownRouter()

	var response *app.GenericResponse
	var expected_response *app.GenericResponse

	var request *http.Request
	var recorder *httptest.ResponseRecorder

	define_route(`GET`, `/users`, server.AllUsers())
	request = initialize_request(`GET`, `/users`, nil)
	recorder = send_request(request)

	// marshall and unmarshall the expected response to be able to check
	// for equality of the data returned.
	expected_byte_response, err := json.Marshal(&app.GenericResponse{
		Status:  true,
		Message: `Users successfully retrieved`,
		Data:    []api.User{userRepo[0], userRepo[1]},
	})

	assert.Equal(t, err, nil)

	json.Unmarshal(expected_byte_response, &expected_response)
	json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expected_response.Status, response.Status)
	assert.Equal(t, expected_response.Message, response.Message)
	assert.Equal(t, expected_response.Data, response.Data)
}

func TestDeleteUser(t *testing.T) {
	SetUpRouter()
	SetUpUserHandlersTest()

	defer TearDownUserHandlersTests()
	defer TearDownRouter()

	var response *app.GenericResponse
	var expected_response *app.GenericResponse

	var request *http.Request
	var recorder *httptest.ResponseRecorder

	define_route(`DELETE`, `/users/:id/delete`, server.DeleteUser())
	request = initialize_request(`DELETE`, `/users/1/delete`, nil)
	recorder = send_request(request)

	expected_byte_response, err := json.Marshal(
		&app.GenericResponse{
			Status:  true,
			Message: `User successfully deleted`,
		},
	)

	assert.Equal(t, err, nil)

	json.Unmarshal(expected_byte_response, &expected_response)
	json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expected_response.Status, response.Status)
	assert.Equal(t, expected_response.Message, response.Message)
}

func TestUpdateUser(t *testing.T) {
	SetUpRouter()
	SetUpUserHandlersTest()

	defer TearDownUserHandlersTests()
	defer TearDownRouter()

	var response *app.GenericResponse
	var expected_response *app.GenericResponse

	var request *http.Request
	var recorder *httptest.ResponseRecorder

	user := api.User{
		Name:  "User One 2.0",
		Email: "user2@email.com",
	}

	jsonValue, _ := json.Marshal(user)

	define_route(`PATCH`, `/users/:id/update`, server.UpdateUser())
	request = initialize_request(`UPDATE`, `/users/1/update`, bytes.NewBuffer(jsonValue))
	recorder = send_request(request)

	expected_response = &app.GenericResponse{
		Status:  true,
		Message: `User successfully updated`,
	}

	json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expected_response, response)
}
