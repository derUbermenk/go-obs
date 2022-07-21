package app_test

import (
	"online-bidding-system/pkg/api"
	"online-bidding-system/pkg/app"
	"os"
	"testing"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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
