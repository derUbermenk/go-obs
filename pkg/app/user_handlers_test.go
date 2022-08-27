package app_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"online-bidding-system/pkg/api"
	"online-bidding-system/pkg/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

var users = map[int]api.User{
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

// mock User Service
type mockUserService struct{}

func (mU *mockUserService) All() ([]api.User, error) {
	// return an array of all the users
	return []api.User{users[0], users[1]}, nil
}

func (mU *mockUserService) Get(userID int) (api.User, error) {
	var user api.User

	user, exists := users[userID]

	if !exists {
		err := &api.ErrNonExistentResource{}
		return user, err
	}

	return user, nil
}

func (mU *mockUserService) Delete(userID int) error {
	_, exists := users[userID]

	if !exists {
		err := &api.ErrNonExistentResource{}
		return err
	}

	return nil
}

func (mU *mockUserService) Update(userID int, user api.User) error {
	/*
		_, exists := mU.userRepo[userID]

		if !exists {
			return &api.ErrNonExistentResource{}
		}
	*/

	return nil
}

// sets up server to use mock user service
// for all user handler tests
func SetUpUserHandlersTest() {
	// initialize api variables
	user_service = &mockUserService{}

	// initialize the server using the initialized services
	server = app.NewServer(
		router, user_service,
		bid_service, bidding_service,
		auth_service,
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
		Data:    []api.User{users[0], users[1]},
	})

	assert.Equal(t, err, nil)

	json.Unmarshal(expected_byte_response, &expected_response)
	json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expected_response.Status, response.Status)
	assert.Equal(t, expected_response.Message, response.Message)
	assert.Equal(t, expected_response.Data, response.Data)
}

func TestGetUser(t *testing.T) {
	SetUpRouter()
	SetUpUserHandlersTest()

	defer TearDownUserHandlersTests()
	defer TearDownRouter()

	var response *app.GenericResponse
	var expected_response *app.GenericResponse

	var request *http.Request
	var recorder *httptest.ResponseRecorder

	define_route(`GET`, `/users/:id`, server.GetUser())

	t.Run(
		"User does not exist",
		func(t *testing.T) {
			request = initialize_request(`GET`, `/users/2`, nil)
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  false,
				Message: "User does not exist",
				Data:    nil,
			}

			json.Unmarshal(recorder.Body.Bytes(), &response)

			assert.Equal(t, http.StatusNotFound, recorder.Code)
			assert.Equal(t, expected_response.Status, response.Status)
			assert.Equal(t, expected_response.Message, response.Message)
			assert.Equal(t, expected_response.Data, response.Data)
		},
	)

	t.Run(
		"User Exists",
		func(t *testing.T) {
			request = initialize_request(`GET`, `/users/1`, nil)
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  true,
				Message: "User retrieved",
				Data:    users[1],
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

	t.Run(
		"User does not exist",
		func(*testing.T) {
			request = initialize_request(`DELETE`, `/users/2/delete`, nil)
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  false,
				Message: `User does not exist`,
			}

			json.Unmarshal(recorder.Body.Bytes(), &response)

			assert.Equal(t, http.StatusBadRequest, recorder.Code)
			assert.Equal(t, expected_response.Status, response.Status)
			assert.Equal(t, expected_response.Message, response.Message)
		},
	)

	t.Run(
		"User exists",
		func(*testing.T) {
			request = initialize_request(`DELETE`, `/users/1/delete`, nil)
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  true,
				Message: `User successfully deleted`,
			}

			json.Unmarshal(recorder.Body.Bytes(), &response)

			assert.Equal(t, http.StatusOK, recorder.Code)
			assert.Equal(t, expected_response.Status, response.Status)
			assert.Equal(t, expected_response.Message, response.Message)
		},
	)
}

/*
func TestUpdateUser(t *testing.T) {
	SetUpRouter()
	SetUpUserHandlersTest()

	defer TearDownUserHandlersTests()
	defer TearDownRouter()

	var response *app.GenericResponse
	var expected_response *app.GenericResponse

	define_route(`PATCH`, `/users/:id/update`, server.UpdateUser())

	var request *http.Request
	var recorder *httptest.ResponseRecorder

	user := api.User{
		Name:  "User One 2.0",
		Email: "user2@email.com",
	}

	// case 1. user does not exist
	t.Run(
		"User does not exist",
		func(t *testing.T) {
			jsonValue, _ := json.Marshal(user)
			request = initialize_request(`PATCH`, `/users/2/update`, bytes.NewBuffer(jsonValue))
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  false,
				Message: `User does not exist`,
			}

			json.Unmarshal(recorder.Body.Bytes(), &response)

			assert.Equal(t, http.StatusBadRequest, recorder.Code)
			assert.Equal(t, expected_response.Data, response.Data)
			assert.Equal(t, expected_response.Message, response.Message)
		},
	)

	t.Run(
		"User exists",
		func(t *testing.T) {
			jsonValue, _ := json.Marshal(user)
			request = initialize_request(`PATCH`, `/users/1/update`, bytes.NewBuffer(jsonValue))
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  true,
				Message: `User successfully updated`,
			}

			json.Unmarshal(recorder.Body.Bytes(), &response)

			assert.Equal(t, http.StatusOK, recorder.Code)
			assert.Equal(t, expected_response.Data, response.Data)
			assert.Equal(t, expected_response.Message, response.Message)
		},
	)
}
*/
