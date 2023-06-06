package bussniss_logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct{}

func (s *Server) GetLogin(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (s *Server) PostLogin(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (s *Server) PostPasswordToken(c *gin.Context, token string) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (s *Server) PostSpace(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (s *Server) DeleteSpaceSpaceId(c *gin.Context, spaceId int) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (s *Server) GetSpaceSpaceId(c *gin.Context, spaceId int) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (s *Server) PutSpaceSpaceId(c *gin.Context, spaceId int) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (s *Server) GetSpaces(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (s *Server) PostUser(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (s *Server) DeleteUserUserId(c *gin.Context, userId int) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (s *Server) GetUserUserId(c *gin.Context, userId int) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (s *Server) PutUserUserId(c *gin.Context, userId int) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (s *Server) GetUsers(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (s *Server) GetWebhooks(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (s *Server) PostWebhooks(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (s *Server) PutWebhooksName(c *gin.Context, name string) {
	c.AbortWithStatus(http.StatusNotImplemented)
}
