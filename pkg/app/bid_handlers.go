package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateBid() gin.HandlerFunc {
	return func(c *gin.Context) {
		createBidRequest := &CreateBidRequest{}

		// get request body
		// parse request body to &api.CreateBidRequest
		err := c.ShouldBindJSON(createBidRequest)
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				&GenericResponse{
					Status:  false,
					Message: err.Error(),
				},
			)

			return
		}

		// call s.bid_service.CreateBid()
		id, err := s.bid_service.CreateBid(
			createBidRequest.BidderID,
			createBidRequest.BiddingID,
			createBidRequest.Amount,
		)

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				&GenericResponse{
					Status:  false,
					Message: err.Error(),
				},
			)
			return
		}

		// return
		c.JSON(
			http.StatusCreated,
			&GenericResponse{
				Status:  true,
				Message: `Bid created`,
				Data:    id,
			},
		)
	}
}
