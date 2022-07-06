package app

import "github.com/gin-gonic/gin"

// initializes the handler funcs for the servers
// router
func (s *Server) Routes() *gin.Engine {
	r := s.router

	// initialize handlers
	v1 := r.Group("v1/api")
	{
		v1.GET("status/", s.ApiStatus())
	}

	return r
}
