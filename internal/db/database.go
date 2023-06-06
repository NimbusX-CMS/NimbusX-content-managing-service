package db

import "github.com/NimbusX-CMS/NimbusX-content-managing-service/internal/models"

type DataBase interface {
	EnsureTablesCreation() error
	GetUser(userId int) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUsers() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(userId int) error
}
