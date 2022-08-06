package app

import (
	"errors"
	"log"
	"net/http"
	"online-bidding-system/pkg/api"
	"strconv"

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

func (s *Server) GetBidding() gin.HandlerFunc {
	return func(c *gin.Context) {
		var id int
		var bidding api.Bidding
		var err error
		id, err = strconv.Atoi(c.Param(`id`))

		if err != nil {
			log.Printf("Handler Error: %v", err)
			c.JSON(
				http.StatusBadRequest,
				&GenericResponse{
					Status:  false,
					Message: "Unable to get bidding",
				},
			)
			return
		}

		bidding, err = s.bidding_service.Get(id)

		if errors.Is(err, &api.ErrNonExistentResource{}) {
			log.Printf("Handler Error: %v", err)
			c.JSON(
				http.StatusNotFound,
				&GenericResponse{
					Status:  false,
					Message: "Bidding does not exist",
				},
			)
			return
		} else if err != nil {
			log.Printf("Handler Error: %v", err)
			c.JSON(
				http.StatusBadRequest,
				&GenericResponse{
					Status:  false,
					Message: "Unable to get bidding",
				},
			)
			return
		}

		c.JSON(
			http.StatusFound,
			&GenericResponse{
				Status:  true,
				Message: "Bidding retrieved",
				Data:    bidding,
			},
		)
	}
}

func (s *Server) UpdateBidding() gin.HandlerFunc {
	return func(c *gin.Context) {
		var id int
		var bidding api.Bidding
		var err error
		id, err = strconv.Atoi(c.Param(`id`))

		if err != nil {
			log.Printf("Handler Error: %v", err)
			c.JSON(
				http.StatusBadRequest,
				&GenericResponse{
					Status:  false,
					Message: "Unable to update bidding",
				},
			)
			return
		}

		err = c.ShouldBindJSON(&bidding)

		if err != nil {
			log.Printf("Handler Error: %v", err)
			c.JSON(
				http.StatusBadRequest,
				&GenericResponse{
					Status:  false,
					Message: "Unable to update bidding",
				},
			)
			return
		}

		err = s.bidding_service.Update(id, bidding)

		if errors.Is(err, &api.ErrNonExistentResource{}) {
			log.Printf("Handler Error: %v", err)
			c.JSON(
				http.StatusBadRequest,
				&GenericResponse{
					Status:  false,
					Message: "Bidding does not exist",
				},
			)
			return
		} else if err != nil {
			log.Printf("Handler Error: %v", err)
			c.JSON(
				http.StatusBadRequest,
				&GenericResponse{
					Status:  false,
					Message: "Unable to update bidding",
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			&GenericResponse{
				Status:  true,
				Message: "Bidding successfully updated",
			},
		)
	}
}

func (s *Server) DeleteBidding() gin.HandlerFunc {
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
					Message: "Unable to delete bidding",
				},
			)
			return
		}

		err = s.bidding_service.Delete(id)

		if errors.Is(err, &api.ErrNonExistentResource{}) {
			log.Printf("Handler Error: %v", err)
			c.JSON(
				http.StatusBadRequest,
				&GenericResponse{
					Status:  false,
					Message: "Bidding does not exist",
				},
			)
			return
		} else if err != nil {
			log.Printf("Handler Error: %v", err)
			c.JSON(
				http.StatusBadRequest,
				&GenericResponse{
					Status:  false,
					Message: "Unable to delete bidding",
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			&GenericResponse{
				Status:  true,
				Message: "Bidding successfully deleted",
			},
		)
	}
}
