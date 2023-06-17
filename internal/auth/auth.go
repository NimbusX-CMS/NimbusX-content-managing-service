package auth

import (
	"fmt"
	"github.com/NimbusX-CMS/NimbusX-content-managing-service/internal/db"
	"github.com/NimbusX-CMS/NimbusX-content-managing-service/internal/error_msg"
	"github.com/NimbusX-CMS/NimbusX-content-managing-service/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const sessionCookieName = "sessionCookie"
const domain = "localhost"

type Auth struct {
	DB db.DataBase
}

func (a Auth) WriteNewCookie(c *gin.Context, emailAndPassword models.EmailAndPassword) bool {
	user, err := a.DB.GetUserByEmail(emailAndPassword.Email)
	if err != nil {
		fmt.Println("Error getting user by email:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return false
	}
	if user == (models.User{}) {
		c.JSON(http.StatusNotFound, error_msg.Error{Error: error_msg.ErrorUserWithEmailNotFound})
		return false
	}
	cookie := models.Session{
		UserID:      user.ID,
		CookieValue: a.generateCookie(),
		ValidUntil:  time.Now().Add(time.Hour * 24 * 30).Unix(), //TODO: make it configurable
	}
	cookie, err = a.DB.CreateSessionCookie(cookie)
	if err != nil {
		fmt.Println("Error creating session cookie:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return false
	}
	c.SetCookie(sessionCookieName, cookie.CookieValue, 60*60*24*30, "", domain, false, true)
	return true
}

func (a Auth) GetSpacePermission(c *gin.Context, spaceId int) (access bool, admin bool) {
	success, session := a.GetSession(c)
	if !success {
		return false, false
	}
	spaceAccess, err := a.DB.GetSpaceAccess(session.UserID, spaceId)
	if err != nil {
		fmt.Println("Error getting space access from DB:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return false, false
	}
	if spaceAccess.SpaceID != spaceId {
		fmt.Println("Error getting space access from DB: spaceId mismatch")
		c.JSON(http.StatusUnauthorized, error_msg.Error{Error: error_msg.ErrorUnauthorizedNoSpaceAccess})
		return false, false
	}
	return true, spaceAccess.Admin
}

func (a Auth) GetOriginPermission(c *gin.Context) bool {
	success, session := a.GetSession(c)
	if !success {
		return false
	}
	return session.User.Origin
}

func (a Auth) GetSession(c *gin.Context) (bool, models.Session) {
	cookie, err := c.Request.Cookie(sessionCookieName)
	if err != nil {
		fmt.Println("Error getting cookie:", err)
		c.JSON(http.StatusUnauthorized, error_msg.Error{Error: error_msg.ErrorUnauthorizedNoSessionCookie})
		return false, models.Session{}
	}
	session, err := a.DB.GetSessionCookieByValue(cookie.Value)
	if err != nil {
		fmt.Println("Error getting session from DB:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return false, models.Session{}
	}
	if session == (models.Session{}) {
		fmt.Println("Session not exist in DB")
		c.JSON(http.StatusUnauthorized, error_msg.Error{Error: error_msg.ErrorUnauthorizedNoSessionNotFound})
		return false, models.Session{}
	}
	if session.ValidUntil < time.Now().Unix() {
		fmt.Println("Session is expired")
		c.JSON(http.StatusUnauthorized, error_msg.Error{Error: error_msg.ErrorUnauthorizedSessionExpired})
		return false, models.Session{}
	}
	return true, session
}

func (a Auth) generateCookie() string {
	return fmt.Sprintf("supperSecretCookieValue-%s", "X") //TODO: make it more secure
}
