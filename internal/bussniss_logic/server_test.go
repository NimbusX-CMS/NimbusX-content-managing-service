package bussniss_logic

import (
	"fmt"
	"github.com/NimbusX-CMS/NimbusX-content-managing-service/internal/models"
	"net/http"
	"testing"
)

func TestPostUser(t *testing.T) {
	w, c, server := setupTest(t)

	tests := TestCases{
		{
			name:               "Create user",
			Url:                "/user",
			RequestBody:        `{"name": "John Doe", "email": "j@example.com"}`,
			ResponseModel:      &models.User{},
			ExpectedBody:       &models.User{ID: 1, Name: "John Doe", Email: "j@example.com"},
			ExpectedStatusCode: http.StatusCreated,
		},
	}

	tests.testStaticUrlCases(t, w, c, server.PostUser)
}

func TestServer_DeleteUserUserId(t *testing.T) {
	w, c, server := setupTest(t)
	user, err := server.DB.CreateUser(models.User{ID: 1, Name: "John Doe", Email: "a@example.com"})
	if err != nil {
		t.Error("Cannot create sample user", err)
	}
	tests := TestCases{
		{
			name:               "Delete user",
			Url:                fmt.Sprintf("/user/%v", user.ID),
			ID:                 user.ID,
			ExpectedStatusCode: http.StatusOK,
		},
	}

	tests.testDynamicIntUrlCases(t, w, c, server.DeleteUserUserId)
}

func TestServer_GetUserUserId(t *testing.T) {
	w, c, server := setupTest(t)
	user, err := server.DB.CreateUser(models.User{ID: 1, Name: "John Doe", Email: "a@example.com"})
	if err != nil {
		t.Error("Cannot create sample user", err)
	}
	tests := TestCases{
		{
			name:               "Get user by id",
			Url:                fmt.Sprintf("/user/%v", user.ID),
			ID:                 user.ID,
			ResponseModel:      &models.User{},
			ExpectedBody:       &models.User{ID: 1, Name: "John Doe", Email: "a@example.com"},
			ExpectedStatusCode: http.StatusOK,
		},
	}

	tests.testDynamicIntUrlCases(t, w, c, server.GetUserUserId)
}

func TestServer_GetUsers(t *testing.T) {
	w, c, server := setupTest(t)
	users := &[]models.User{
		{ID: 1, Name: "John Doe", Email: "a@example.com"},
		{ID: 2, Name: "Jane Doe", Email: "b@example.com"},
	}

	for _, user := range *users {
		_, err := server.DB.CreateUser(user)
		if err != nil {
			t.Error("Cannot create sample user", err)
		}
	}

	tests := TestCases{
		{
			name:               "Get users",
			Url:                "/users",
			ResponseModel:      &[]models.User{},
			ExpectedBody:       users,
			ExpectedStatusCode: http.StatusOK,
		},
	}

	tests.testStaticUrlCases(t, w, c, server.GetUsers)
}

func TestServer_PutUserUserId(t *testing.T) {
	w, c, server := setupTest(t)
	user, err := server.DB.CreateUser(models.User{ID: 1, Name: "John Doee", Email: "a@example.com"})
	if err != nil {
		t.Error("Cannot create sample user", err)
	}
	tests := TestCases{
		{
			name:               "Put user by id",
			Url:                fmt.Sprintf("/user/%v", user.ID),
			ID:                 user.ID,
			RequestBody:        `{"name": "John Doe", "email": "a@example.com"}`,
			ResponseModel:      &models.User{},
			ExpectedBody:       &models.User{ID: 1, Name: "John Doe", Email: "a@example.com"},
			ExpectedStatusCode: http.StatusOK,
		},
	}

	tests.testDynamicIntUrlCases(t, w, c, server.PutUserUserId)
}
