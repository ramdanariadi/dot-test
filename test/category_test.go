package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/assert/v2"
	"github.com/ramdanariadi/dot-test/category"
	"github.com/ramdanariadi/dot-test/helpers"
	"github.com/ramdanariadi/dot-test/route"
	"github.com/ramdanariadi/dot-test/setup"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var connection, _ = setup.NewDbConnection() // can be replaced with different db connection such as sqlite, mysql, postgres, sqlserver.
var db, _ = gorm.Open(postgres.New(postgres.Config{Conn: connection}))
var router = route.RouterSetup(db)

var categories []*category.Category

func authorization() string {
	recorder := httptest.NewRecorder()
	body := []byte(`{"username" : "admin","password" : "password"}`)
	request, _ := http.NewRequest("POST", "/login", bytes.NewReader(body))
	request.Header.Add("Content-Type", "Application/json")
	router.ServeHTTP(recorder, request)

	var tokens map[string]any
	bodyBytes, _ := io.ReadAll(recorder.Body)
	err := json.Unmarshal(bodyBytes, &tokens)
	helpers.LogIfError(err)

	return fmt.Sprintf("Bearer %s", tokens["accessToken"].(string))
}

func TestCategory(t *testing.T) {

	t.Run("store category", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		body := []byte(`{"category": "Fish"}`)
		request, _ := http.NewRequest("POST", "/category/", bytes.NewBuffer(body))
		request.Header.Add("Authorization", authorization())
		router.ServeHTTP(recorder, request)
		assert.Equal(t, 200, recorder.Code)
	})

	t.Run("get all category", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/category/", nil)
		router.ServeHTTP(recorder, request)

		bytes, _ := io.ReadAll(recorder.Body)
		err := json.Unmarshal(bytes, &categories)
		assert.Equal(t, err, nil)
		assert.Equal(t, 200, recorder.Code)
	})

	t.Run("update category", func(t *testing.T) {
		if len(categories) > 0 {
			categoryToUpdate := categories[0]
			recorder := httptest.NewRecorder()
			body := []byte(`{"category": "Fish"}`)
			request, _ := http.NewRequest("PUT", "/category/"+categoryToUpdate.ID, bytes.NewBuffer(body))
			request.Header.Add("Authorization", authorization())
			router.ServeHTTP(recorder, request)
			assert.Equal(t, 200, recorder.Code)
		}
	})
}
