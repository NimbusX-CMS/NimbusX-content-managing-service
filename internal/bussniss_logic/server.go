package bussniss_logic

import (
	"fmt"
	"github.com/NimbusX-CMS/NimbusX-content-managing-service/internal/db"
	"github.com/NimbusX-CMS/NimbusX-content-managing-service/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	DB db.DataBase
}

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
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := s.DB.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (s *Server) DeleteUserUserId(c *gin.Context, userId int) {
	err := s.DB.DeleteUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (s *Server) GetUserUserId(c *gin.Context, userId int) {
	user, err := s.DB.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (s *Server) PutUserUserId(c *gin.Context, userId int) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = userId

	updatedUser, err := s.DB.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("updatedUser", updatedUser)
	c.JSON(http.StatusOK, updatedUser)
}

func (s *Server) GetUsers(c *gin.Context) {
	users, err := s.DB.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (s *Server) GetUserUserIdSpaces(c *gin.Context, userId int) {
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
