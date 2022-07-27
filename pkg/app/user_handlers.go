package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) AllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := s.user_service.All()

		if err != nil {
			log.Printf("Service Error: %v", err)
			c.JSON(
				http.StatusInternalServerError,
				&GenericResponse{
					Status:  false,
					Message: "Error retrieving users",
				},
			)
		}

		c.JSON(
			http.StatusOK,
			&GenericResponse{
				Status:  true,
				Message: "Users successfully retrieved",
				Data:    users,
			},
		)
	}
}
