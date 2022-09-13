package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) LogIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginRequest := &LoginRequest{}

		// get the request body
		err := c.ShouldBindJSON(loginRequest)
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

		// login the user
		err, accessToken, refreshToken := s.auth_service.LogIn(loginRequest.Email, loginRequest.Password)

		if err != nil {
			if err.Error() == `Invalid credentials` {
				c.JSON(
					http.StatusUnauthorized,
					&GenericResponse{
						Status:  false,
						Message: `Invalid credentials`,
					},
				)

				return

			} else {
				c.JSON(
					http.StatusInternalServerError,
					&GenericResponse{
						Status:  false,
						Message: ``,
					},
				)

				return
			}
		}

		c.JSON(
			http.StatusOK,
			&GenericResponse{
				Status:  true,
				Message: `Logged in`,
				Data: &AuthResponse{
					AccessToken:  accessToken,
					RefreshToken: refreshToken,
				},
			},
		)
	}
}
