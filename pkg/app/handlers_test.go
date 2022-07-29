package app_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"online-bidding-system/pkg/api"
	"online-bidding-system/pkg/app"
	"testing"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

var server *app.Server
var router *gin.Engine

// initialize api variables
var user_service api.UserService
var bidding_service api.BiddingService
var auth_service api.AuthService

// sets up the router. Call this at the start of every test
// to ensure router defintion is independent of all tests.
//
// had problems earlier where any endpoint with params can not read
// the params when defined after an endpoint without any params.
// this was solved by either:
//
// 1. defining a dummy endpoint at the start
// of all tests (problematic as requires to keep track of test order)
// or
// 2. define router setup to be run for every test which led to this function
func SetUpRouter() {
	router = gin.Default()
	router.Use(cors.Default())
}

// tears down the router, effectively setting it to nil for formality's sake.
func TearDownRouter() {
	router = nil
}

// defines the route to be used for testing the handler
func define_route(method string, path string, handler_func gin.HandlerFunc) {
	router.Handle(method, path, handler_func)
}

func initialize_request(method string, path string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, path, body)
	return req
}

func send_request(req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func TestApiStatus(t *testing.T) {
	SetUpRouter()
	defer TearDownRouter()

	server = app.NewServer(router, user_service, bidding_service, auth_service)

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
