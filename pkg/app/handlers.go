package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenericResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, &GenericResponse{
			Status:  true,
			Message: "Bidding System API running smoothly",
		})
	}
}
