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

var server *app.Server
var router *gin.Engine

func TestMain(m *testing.M) {
	// initialize api variables
	var user_service api.UserService
	var bidding_service api.BiddingService
	var auth_service api.AuthService

	// initialize a router
	router = gin.Default()
	router.Use(cors.Default())
	server = app.NewServer(router, user_service, bidding_service, auth_service)
	// initialize a server with the router

	exitValue := m.Run()
	os.Exit(exitValue)
}

func TestApiStatus(t *testing.T) {
	// define the route
	router.GET(`/v1/api/status`, server.ApiStatus())
	req, _ := http.NewRequest(`GET`, `/v1/api/status`, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response *app.GenericResponse
	expected_response := &app.GenericResponse{
		Status:  true,
		Message: "Bidding System API running smoothly",
	}

	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected_response, response)
}
