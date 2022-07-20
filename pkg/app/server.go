package app

import (
	"log"
	"online-bidding-system/pkg/api"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router          *gin.Engine
	user_service    api.UserService
	bidding_service api.BiddingService
	auth_service    api.AuthService
}

func NewServer(
	router *gin.Engine, user_service api.UserService,
	bidding_service api.BiddingService, auth_service api.AuthService,
) *Server {
	return &Server{
		router:          router,
		user_service:    user_service,
		bidding_service: bidding_service,
		auth_service:    auth_service,
	}
}

func (s *Server) Run() error {
	// run function that initialisez the routes
	r := s.Routes()

	// run the server through the router
	err := r.Run()

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
