package app

import (
	"errors"
	"log"
	"net/http"
	"online-bidding-system/pkg/api"
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

			return
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

func (s *Server) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var id int
		var user api.User
		var err error
		id, err = strconv.Atoi(c.Param(`id`))

		if err != nil {
			log.Printf("Handler Error: %v", err)
			c.JSON(
				http.StatusBadRequest,
				&GenericResponse{
					Status:  false,
					Message: "Unable to get user",
				},
			)
			return
		}

		user, err = s.user_service.Get(id)

		if errors.Is(err, &api.ErrNonExistentUser{}) {
			log.Printf("Handler Error: %v", err)
			c.JSON(
				http.StatusNotFound,
				&GenericResponse{
					Status:  false,
					Message: "User does not exist",
				},
			)
			return
		} else if err != nil {
			log.Printf("Handler Error: %v", err)
			c.JSON(
				http.StatusBadRequest,
				&GenericResponse{
					Status:  false,
					Message: "Unable to get user",
				},
			)
			return
		}

		c.JSON(
			http.StatusFound,
			&GenericResponse{
				Status:  true,
				Message: "User retrieved",
				Data:    user,
			},
		)
	}
}

func (s *Server) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var id int
		var err error
		id, err = strconv.Atoi(c.Param(`id`))

		if err != nil {
			log.Printf("Handler Error: %v", err)
			c.JSON(
				http.StatusBadRequest,
				&GenericResponse{
					Status:  false,
					Message: "Unable to update user",
				},
			)
			return
		}

		err = s.user_service.Delete(id)

		if errors.Is(err, &api.ErrNonExistentUser{}) {
			log.Printf("Handler Error: %v", err)
			c.JSON(
				http.StatusBadRequest,
				&GenericResponse{
					Status:  false,
					Message: "User does not exist",
				},
			)
			return
		} else if err != nil {
			log.Printf("Handler Error: %v", err)
			c.JSON(
				http.StatusBadRequest,
				&GenericResponse{
					Status:  false,
					Message: "Unable to delete user",
				},
			)
			return
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

func (s *Server) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var id int
		var user api.User
		var err error
		id, err = strconv.Atoi(c.Param(`id`))

		if err != nil {
			log.Printf("Handler Error: %v", err)
			c.JSON(
				http.StatusBadRequest,
				&GenericResponse{
					Status:  false,
					Message: "Unable to update user",
				},
			)
			return
		}

		err = c.ShouldBindJSON(&user)

		if err != nil {
			log.Printf("Handler Error: %v", err)
			c.JSON(
				http.StatusBadRequest,
				&GenericResponse{
					Status:  false,
					Message: "Unable to update user",
				},
			)
			return
		}

		err = s.user_service.Update(id, user)

		if errors.Is(err, &api.ErrNonExistentUser{}) {
			log.Printf("Handler Error: %v", err)
			c.JSON(
				http.StatusBadRequest,
				&GenericResponse{
					Status:  false,
					Message: "User does not exist",
				},
			)
			return
		} else if err != nil {
			log.Printf("Handler Error: %v", err)
			c.JSON(
				http.StatusBadRequest,
				&GenericResponse{
					Status:  false,
					Message: "Unable to update user",
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			&GenericResponse{
				Status:  true,
				Message: "User successfully updated",
			},
		)
	}
}
