package multi_db

import (
	"errors"
	"github.com/NimbusX-CMS/NimbusX-content-managing-service/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MultiDB struct {
	db *gorm.DB
}

func NewMultiDB(db *gorm.DB) *MultiDB {
	return &MultiDB{
		db: db,
	}
}

func (m *MultiDB) EnsureTablesCreation() error {
	return m.db.AutoMigrate(&models.User{})
}

func (m *MultiDB) GetUser(userId int) (models.User, error) {
	var user models.User
	err := m.db.First(&user, userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, nil
		}
		return models.User{}, err
	}
	return user, err
}

func (m *MultiDB) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := m.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, nil
		}
		return models.User{}, err
	}
	return user, nil
}

func (m *MultiDB) GetUsers() ([]models.User, error) {
	var users []models.User
	err := m.db.Find(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []models.User{}, nil
		}
		return []models.User{}, err
	}

	return users, err
}

func (m *MultiDB) CreateUser(user models.User) (models.User, error) {
	user.ID = 0
	err := m.db.Create(&user).Error
	return user, err
}

func (m *MultiDB) UpdateUser(user models.User) (models.User, error) {
	err := m.db.Save(&user).Error
	return user, err
}

func (m *MultiDB) DeleteUser(userId int) error {
	return m.db.Delete(&models.User{}, userId).Error
}

func ConnectToSQLite(databasePath string) (*MultiDB, error) {
	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return NewMultiDB(db), nil
}
