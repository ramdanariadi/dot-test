package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/ramdanariadi/dot-test/setup"
	"net/http/httptest"
	"testing"
)

func TestAddProduct(t *testing.T) {
	router := setup.SetupRouter()
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/category", nil)
	router.ServeHTTP(recorder, request)
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, gin.H{}, recorder.Body.String())
}
