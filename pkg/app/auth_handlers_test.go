package app_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"online-bidding-system/pkg/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockAuthService struct{}

// initializes server for this test using the
// auth service we are using here
func SetUpAuthHandlersTest() {
	auth_service = &mockAuthService{}

	server = app.NewServer(
		router, user_service,
		bid_service, bidding_service,
		auth_service,
	)
}

// tear down the server we are using, resets the
// auth_service and server to nil
func TearDownAuthHandlersTests() {
	auth_service = nil
	server = nil
}

func TestLogIn(t *testing.T) {
	SetUpRouter()
	SetUpAuthHandlersTest()

	defer TearDownRouter()
	defer TearDownAuthHandlersTests()

	var response *app.GenericResponse
	var expected_response *app.GenericResponse

	var request *http.Request
	var recorder *httptest.ResponseRecorder

	define_route(`GET`, `/authentication/login`, server.LogIn())

	t.Run(
		`Correct credentials`,
		func(t *testing.T) {
			loginRequest := &app.LoginRequest{
				Email:    `existingUser@email.com`,
				Password: `correctPassword`,
			}

			jsonValue, _ := json.Marshal(loginRequest)

			request = initialize_request(``, `/authentication/login`, bytes.NewBuffer(jsonValue))
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  true,
				Message: `Logged in`,
				Data: &app.AuthResponse{
					AccessToken:  `ValidAccessToken`,
					RefreshToken: `ValidRefreshToken`,
				},
			}

			json.Unmarshal(recorder.Body.Bytes(), &response)

			assert.Equal(t, http.StatusOK, recorder.Code)
			assert.Equal(t, expected_response, response.Status)
			assert.Equal(t, expected_response.Message, response.Message)
			assert.Equal(
				t,
				expected_response.Data.(*app.AuthResponse).AccessToken,
				response.Data.(*app.AuthResponse).AccessToken,
			)
			assert.Equal(
				t,
				expected_response.Data.(*app.AuthResponse).RefreshToken,
				response.Data.(*app.AuthResponse).RefreshToken,
			)
		},
	)

	t.Run(
		`Incorrect credentials`,
		func(t *testing.T) {
			loginRequest := &app.LoginRequest{
				Email:    `nonexistingUser@email.com`,
				Password: `incorrectPassword`,
			}

			jsonValue, _ := json.Marshal(loginRequest)

			request = initialize_request(``, `/authentication/login`, bytes.NewBuffer(jsonValue))
			recorder = send_request(request)

			expected_response = &app.GenericResponse{
				Status:  true,
				Message: `Logged in`,
				Data:    nil,
			}

			json.Unmarshal(recorder.Body.Bytes(), &response)

			assert.Equal(t, http.StatusOK, recorder.Code)
			assert.Equal(t, expected_response, response.Status)
			assert.Equal(t, expected_response.Message, response.Message)
			assert.Equal(t, expected_response.Data, response.Data)
		},
	)
}