package db

import "github.com/NimbusX-CMS/NimbusX-content-managing-service/internal/models"

type DataBase interface {
	EnsureTablesCreation() error

	GetSessionCookieByValue(cookieValue string) (models.Session, error)
	GetSessionCookiesByUserId(userId int) ([]models.Session, error)
	CreateSessionCookie(cookie models.Session) (models.Session, error)
	UpdateSessionCookie(cookie models.Session) (models.Session, error)
	DeleteSessionCookie(cookieId int) error

	GetUser(userId int) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUsers() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(userId int) error

	GetSpace(spaceId int) (models.Space, error)
	GetSpaces() ([]models.Space, error)
	CreateSpace(space models.Space) (models.Space, error)
	UpdateSpace(space models.Space) (models.Space, error)
	DeleteSpace(spaceId int) error

	GetSpaceAccess(userId int, spaceId int) (models.SpaceAccess, error)
	GetSpaceAccesses(userId int) ([]models.SpaceAccess, error)
	CreateSpaceAccess(spaceAccess models.SpaceAccess) (models.SpaceAccess, error)
	UpdateSpaceAccess(spaceAccess models.SpaceAccess) (models.SpaceAccess, error)
	DeleteSpaceAccess(userId int, spaceId int) error
}
