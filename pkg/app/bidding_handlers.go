package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) AllBiddings() gin.HandlerFunc {
	return func(c *gin.Context) {
		biddings, err := s.bidding_service.All()

		if err != nil {
			log.Printf("Service Error: %v", err)
			c.JSON(
				http.StatusInternalServerError,
				&GenericResponse{
					Status:  false,
					Message: "Error retrieving biddings",
				},
			)

			return
		}

		c.JSON(
			http.StatusOK,
			&GenericResponse{
				Status:  true,
				Message: "Biddings successfully retrieved",
				Data:    biddings,
			},
		)
	}
}
