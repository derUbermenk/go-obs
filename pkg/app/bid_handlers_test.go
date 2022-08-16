package app_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"online-bidding-system/pkg/app"
	"testing"

	"github.com/go-playground/assert/v2"
)

func SetUpBidHandlersTest() {

}

func TearDownBidHandlersTest() {

}

func TestCreateBid(t *testing.T) {
	SetUpRouter()
	SetUpBidHandlersTest()

	defer TearDownBidHandlersTest()
	defer TearDownRouter()

	var response *app.GenericResponse
	var expected_response *app.GenericResponse

	var request *http.Request
	var recorder *httptest.ResponseRecorder

	define_route(`POST`, `/bids`, server.CreateBid())

	request = initialize_request(`POST`, `/bids`, nil)
	recorder = send_request(request)

	json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Equal(t, http.StatusOK, recorder.Code)
}
