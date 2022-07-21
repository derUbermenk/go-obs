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

type User struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashedPass"`
}

var server *app.Server
var router *gin.Engine

func TestMain(m *testing.M) {
	// initialize api variables
	var user_service api.UserService
	var bidding_service api.BiddingService
	var auth_service api.AuthService

	// initialize router
	router = gin.Default()
	router.Use(cors.Default())
	server = app.NewServer(router, user_service, bidding_service, auth_service)

	// initialize a server with the router

	// run the tests
	exitValue := m.Run()
	os.Exit(exitValue)
}

func TestAllUsers(t *testing.T) {
	router.GET(`/v1/api/status`, server.AllUsers())
	req, _ := http.NewRequest(`GET`, `/v1/api/status`, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	user1 := User{
		Name:           "User One",
		Email:          "user1@email.com",
		HashedPassword: "x120asd",
	}

	user2 := User{
		Name:           "User Two",
		Email:          "user2@email.com",
		HashedPassword: "y562ash",
	}

	var response *app.GenericResponse
	expected_response := &app.GenericResponse{
		Status:  true,
		Message: `Users successfully retrieved`,
		Data:    []User{user1, user2},
	}

	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected_response, response)
}
