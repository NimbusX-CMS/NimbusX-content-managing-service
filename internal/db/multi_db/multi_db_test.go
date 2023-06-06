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

	resultUser, err := multiDB.GetUser(createdUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, createdUser.ID, resultUser.ID)
	assert.Equal(t, createdUser.Name, resultUser.Name)
	assert.Equal(t, createdUser.Email, resultUser.Email)
}

func TestMultiDB_GetUserByEmail(t *testing.T) {
	multiDB, teardown := setupTestDB()
	defer teardown()

	user := models.User{
		Name:  "Jane Smith",
		Email: "jane@example.com",
	}

	_, _ = multiDB.CreateUser(user)

	resultUser, err := multiDB.GetUserByEmail(user.Email)
	assert.NoError(t, err)
	assert.Equal(t, user.Name, resultUser.Name)
	assert.Equal(t, user.Email, resultUser.Email)
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

	resultUsers, err := multiDB.GetUsers()
	assert.NoError(t, err)
	assert.Equal(t, len(users), len(resultUsers))

	for i, user := range users {
		assert.Equal(t, user.Name, resultUsers[i].Name)
		assert.Equal(t, user.Email, resultUsers[i].Email)
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

	createdUser.Name = "John Smith"
	updatedUser, err := multiDB.UpdateUser(createdUser)
	assert.NoError(t, err)
	assert.Equal(t, createdUser.ID, updatedUser.ID)
	assert.Equal(t, createdUser.Name, updatedUser.Name)
	assert.Equal(t, createdUser.Email, updatedUser.Email)
}

func TestMultiDB_DeleteUser(t *testing.T) {
	multiDB, teardown := setupTestDB()
	defer teardown()

	user := models.User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	createdUser, _ := multiDB.CreateUser(user)

	err := multiDB.DeleteUser(createdUser.ID)
	assert.NoError(t, err)

	_, err = multiDB.GetUser(createdUser.ID)
	assert.Error(t, err)
}
