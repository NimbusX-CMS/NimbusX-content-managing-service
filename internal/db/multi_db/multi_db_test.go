package multi_db

import (
	"testing"

	"github.com/NimbusX-CMS/NimbusX-content-managing-service/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*MultiDB, func()) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	multiDB := NewMultiDB(db)

	err := multiDB.EnsureTablesCreation()
	if err != nil {
		panic("Failed to migrate database")
	}

	return multiDB, func() {
		_ = db.Migrator().DropTable(&models.User{})
	}
}

func TestMultiDB_GetUser(t *testing.T) {
	multiDB, teardown := setupTestDB()
	defer teardown()

	user := models.User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	createdUser, _ := multiDB.CreateUser(user)

	tests := []struct {
		name          string
		id            int
		expectedUser  models.User
		expectedError error
	}{
		{
			name:          "Existing User",
			id:            createdUser.ID,
			expectedUser:  models.User{ID: createdUser.ID, Name: "John Doe", Email: "john@example.com"},
			expectedError: nil,
		},
		{
			name:          "Non-Existing User",
			id:            1234,
			expectedUser:  models.User{},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resultUser, err := multiDB.GetUser(test.id)
			assert.Equal(t, test.expectedUser, resultUser)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestMultiDB_GetUserByEmail(t *testing.T) {
	multiDB, teardown := setupTestDB()
	defer teardown()

	user := models.User{
		Name:  "John",
		Email: "existing@example.com",
	}

	_, _ = multiDB.CreateUser(user)

	tests := []struct {
		name         string
		email        string
		expectedUser models.User
	}{
		{
			name:         "Existing User",
			email:        "existing@example.com",
			expectedUser: models.User{ID: 1, Name: "John", Email: "existing@example.com"},
		},
		{
			name:         "Non-Existing User",
			email:        "not-existing@example.com",
			expectedUser: models.User{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resultUser, err := multiDB.GetUserByEmail(test.email)
			assert.Equal(t, test.expectedUser, resultUser)
			assert.Equal(t, nil, err)
		})
	}
}

func TestMultiDB_GetUsers(t *testing.T) {
	multiDB, teardown := setupTestDB()
	defer teardown()

	users := []models.User{
		{
			Name:  "Alice",
			Email: "alice@example.com",
		},
		{
			Name:  "Bob",
			Email: "bob@example.com",
		},
		{
			Name:  "Charlie",
			Email: "charlie@example.com",
		},
	}

	for _, user := range users {
		_, _ = multiDB.CreateUser(user)
	}

	tests := []struct {
		name           string
		expectedLength int
	}{
		{
			name:           "Get All Users",
			expectedLength: len(users),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resultUsers, err := multiDB.GetUsers()
			assert.NoError(t, err)
			assert.Len(t, resultUsers, test.expectedLength)

			for i, user := range users {
				assert.Equal(t, user.Name, resultUsers[i].Name)
				assert.Equal(t, user.Email, resultUsers[i].Email)
			}
		})
	}
}

func TestMultiDB_UpdateUser(t *testing.T) {
	multiDB, teardown := setupTestDB()
	defer teardown()

	user := models.User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	createdUser, _ := multiDB.CreateUser(user)

	tests := []struct {
		name         string
		userToUpdate models.User
		expectedUser models.User
		expectedErr  error
	}{
		{
			name:         "Update User",
			userToUpdate: models.User{ID: createdUser.ID, Name: "John Smith", Email: "john@example.com"},
			expectedUser: models.User{ID: createdUser.ID, Name: "John Smith", Email: "john@example.com"},
			expectedErr:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			updatedUser, err := multiDB.UpdateUser(test.userToUpdate)
			assert.NoError(t, err)
			assert.Equal(t, test.expectedUser, updatedUser)
		})
	}
}

func TestMultiDB_DeleteUser(t *testing.T) {
	multiDB, teardown := setupTestDB()
	defer teardown()

	user := models.User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	createdUser, _ := multiDB.CreateUser(user)

	tests := []struct {
		name         string
		idToDelete   int
		expectedErr  error
		expectedUser models.User
	}{
		{
			name:         "Delete User",
			idToDelete:   createdUser.ID,
			expectedErr:  nil,
			expectedUser: models.User{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := multiDB.DeleteUser(test.idToDelete)
			assert.NoError(t, err)

			deletedUser, err := multiDB.GetUser(test.idToDelete)
			assert.NoError(t, err)
			assert.Equal(t, test.expectedUser, deletedUser)
		})
	}
}
