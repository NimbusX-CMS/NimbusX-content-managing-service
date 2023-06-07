package bussniss_logic

import (
	"encoding/json"
	"fmt"
	"github.com/NimbusX-CMS/NimbusX-content-managing-service/internal/db"
	"github.com/NimbusX-CMS/NimbusX-content-managing-service/internal/db/multi_db"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type TestCases []TestCase

type TestCase struct {
	name               string
	Url                string
	ID                 int
	Str                string
	RequestBody        string
	ResponseModel      any
	ExpectedBody       interface{}
	ExpectedStatusCode int
}

func (tc TestCases) testStaticUrlCases(t *testing.T, w *httptest.ResponseRecorder, c *gin.Context, toTest func(ctx *gin.Context)) {
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			c.Request, _ = http.NewRequest(http.MethodPost, tt.Url, strings.NewReader(tt.RequestBody))
			c.Request.Header.Set("Content-Type", "application/json")

			toTest(c)

			AssertStatusCode(t, w, tt.ExpectedStatusCode)

			AssertBody(t, w, tt.ResponseModel, tt.ExpectedBody)
		})
	}
}
func (tc TestCases) testDynamicIntUrlCases(t *testing.T, w *httptest.ResponseRecorder, c *gin.Context, toTest func(ctx *gin.Context, id int)) {
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			c.Request, _ = http.NewRequest(http.MethodPost, tt.Url, strings.NewReader(tt.RequestBody))
			c.Request.Header.Set("Content-Type", "application/json")

			toTest(c, tt.ID)

			AssertStatusCode(t, w, tt.ExpectedStatusCode)

			AssertBody(t, w, tt.ResponseModel, tt.ExpectedBody)
		})
	}
}

func (tc TestCases) testDynamicStringUrlCases(t *testing.T, w *httptest.ResponseRecorder, c *gin.Context, toTest func(ctx *gin.Context, str string)) {
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			c.Request, _ = http.NewRequest(http.MethodPost, tt.Url, strings.NewReader(tt.RequestBody))
			c.Request.Header.Set("Content-Type", "application/json")

			toTest(c, tt.Str)

			AssertStatusCode(t, w, tt.ExpectedStatusCode)

			AssertBody(t, w, tt.ResponseModel, tt.ExpectedBody)
		})
	}
}

func setupTestDB(t *testing.T) db.DataBase {
	testDB, err := multi_db.ConnectToSQLite(":memory:")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		t.Error("Error connecting to database:", err)
	}
	err = testDB.EnsureTablesCreation()
	if err != nil {
		fmt.Println("Error creating tables:", err)
		t.Error("Error connecting to database:", err)
	}
	return testDB
}

func setupTest(t *testing.T) (*httptest.ResponseRecorder, *gin.Context, *Server) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	testDB := setupTestDB(t)
	server := &Server{DB: testDB}

	return w, c, server
}

func AssertStatusCode(t *testing.T, w *httptest.ResponseRecorder, expectedCode int) {
	if w.Code != expectedCode {
		t.Errorf("expected status %v, got %v", expectedCode, w.Code)
	}
}

func AssertBody(t *testing.T, w *httptest.ResponseRecorder, responseModel any, expectedBody interface{}) {
	if responseModel == nil {
		return
	}
	var err = json.Unmarshal(w.Body.Bytes(), responseModel)
	if err != nil {
		t.Errorf("Error parsing response body: %v\n%v", err, w.Body.String())
	}

	if !reflect.DeepEqual(responseModel, expectedBody) {
		t.Errorf("Expected user %+v but got %+v", expectedBody, responseModel)
	}
}
