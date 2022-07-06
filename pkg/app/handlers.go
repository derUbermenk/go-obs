package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := map[string]string{
			"status": "success",
			"data":   "obs api running smoothly",
		}

		c.JSON(http.StatusOK, response)
	}
}
