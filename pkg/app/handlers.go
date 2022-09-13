package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// reponse

type GenericResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AuthResponse struct {
	AccessToken  string `json:"status"`
	RefreshToken string `json:"message"`
}

// request

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, &GenericResponse{
			Status:  true,
			Message: "Bidding System API running smoothly",
		})
	}
}
