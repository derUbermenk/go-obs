package app_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"online-bidding-system/pkg/api"
	"online-bidding-system/pkg/app"
	"os"
	"testing"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

// this must be the same as that defined in api definition of the user.
// this is made so that no api code is called in this test
type mockUserService struct {
	userRepo map[int]api.User
}

var userRepo = map[int]api.User{
	0: api.User{
		Name:           "User One",
		Email:          "user1@email.com",
		HashedPassword: "x120asd",
	},
	1: api.User{
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
	return nil
}

func (mU *mockUserService) Update(user api.User) error {
	return nil
}

var server *app.Server
var router *gin.Engine

func TestMain(m *testing.M) {
	// initialize router
	router = gin.Default()
	router.Use(cors.Default())

	// run the tests
	exitValue := m.Run()
	os.Exit(exitValue)
}

func TestAllUsers(t *testing.T) {
	// initialize api variables
	var user_service *mockUserService
	var bidding_service api.BiddingService
	var auth_service api.AuthService
	user_service = &mockUserService{userRepo: userRepo}

	// initialize the server using the initialized services
	server = app.NewServer(router, user_service, bidding_service, auth_service)

	router.GET(`/v1/api/status`, server.AllUsers())
	req, _ := http.NewRequest(`GET`, `/v1/api/status`, nil)
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
	var user_service *mockUserService
	var bidding_service api.BiddingService
	var auth_service api.AuthService
	user_service = &mockUserService{userRepo: userRepo}

	// initialize the server using the initialized services
	server = app.NewServer(router, user_service, bidding_service, auth_service)

	// prepare the request
	router.GET(`/v1/api/user/1/delete`, server.DeleteUser())
	req, _ := http.NewRequest(`DELETE`, `/v1/api/user/1/delete`, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response *app.GenericResponse
	var expected_response *app.GenericResponse

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
