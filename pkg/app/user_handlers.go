package app

import (
	"log"
	"net/http"
	"strconv"

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

func (s *Server) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		delete_userID, err := strconv.Atoi(c.Param(`id`))

		if err != nil {
			log.Printf("Handler Error: %v", err)
			c.JSON(
				http.StatusInternalServerError,
				&GenericResponse{
					Status:  false,
					Message: "Error handling Request",
				},
			)
		}

		err = s.user_service.Delete(delete_userID)

		if err != nil {
			log.Printf("Service Error: %v", err)
			c.JSON(
				http.StatusInternalServerError,
				&GenericResponse{
					Status:  false,
					Message: "Error deleting user",
				},
			)
		}

		c.JSON(
			http.StatusOK,
			&GenericResponse{
				Status:  true,
				Message: "User successfully deleted",
			},
		)
	}
}
