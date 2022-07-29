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

	router.GET(`/v1/api/users`, server.AllUsers())
	req, _ := http.NewRequest(`GET`, `/v1/api/users`, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response *app.GenericResponse
	var expected_response *app.GenericResponse

	// marshall and unmarshall the expected response to be able to check
	// for equality of the data returned.
	expected_byte_response, err := json.Marshal(&app.GenericResponse{
		Status:  true,
		Message: `Users successfully retrieved`,
		Data:    []api.User{userRepo[0], userRepo[1]},
	})

	assert.Equal(t, err, nil)

	json.Unmarshal(expected_byte_response, &expected_response)
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected_response.Status, response.Status)
	assert.Equal(t, expected_response.Message, response.Message)
	assert.Equal(t, expected_response.Data, response.Data)
}

func TestDeleteUser(t *testing.T) {
	SetUpRouter()
	SetUpUserHandlersTest()

	defer TearDownUserHandlersTests()
	defer TearDownRouter()

	// prepare the request
	router.DELETE("/user/:id/delete", server.DeleteUser())
	req, _ := http.NewRequest("DELETE", "/user/1/delete", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	response := &app.GenericResponse{}
	expected_response := &app.GenericResponse{}

	expected_byte_response, err := json.Marshal(
		&app.GenericResponse{
			Status:  true,
			Message: `User successfully deleted`,
		},
	)

	assert.Equal(t, err, nil)

	json.Unmarshal(expected_byte_response, expected_response)
	json.Unmarshal(w.Body.Bytes(), response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected_response.Status, response.Status)
	assert.Equal(t, expected_response.Message, response.Message)
}
