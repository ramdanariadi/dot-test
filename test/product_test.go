package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/assert/v2"
	"github.com/ramdanariadi/dot-test/product"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var products []*product.Product

func TestProduct(t *testing.T) {
	t.Run("store product", func(t *testing.T) {
		if len(categories) > 0 {
			recorder := httptest.NewRecorder()
			category := categories[0]
			body := []byte(fmt.Sprintf(`{
		  "name": "Octopus",
		  "price": 10,
		  "weight": 1000,
		  "description": "Good for health",
		  "imageUrl" : null,
		  "categoryId": "%s"
		}`, category.ID))
			request, _ := http.NewRequest("POST", "/product/", bytes.NewBuffer(body))
			request.Header.Add("Authorization", authorization())
			router.ServeHTTP(recorder, request)
			assert.Equal(t, 200, recorder.Code)
		}
	})

	t.Run("get all products", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/product/", nil)
		router.ServeHTTP(recorder, request)

		bodyBytes, _ := io.ReadAll(recorder.Body)
		err := json.Unmarshal(bodyBytes, &products)
		assert.Equal(t, err, nil)
		assert.Equal(t, 200, recorder.Code)
	})

	t.Run("get product by id", func(t *testing.T) {
		if len(products) > 0 {
			p := products[0]
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/product/"+p.ID, nil)
			router.ServeHTTP(recorder, request)

			var productById product.Product
			productBytes, _ := io.ReadAll(recorder.Body)
			err := json.Unmarshal(productBytes, &productById)
			assert.Equal(t, err, nil)
			assert.Equal(t, 200, recorder.Code)
		}
	})

	t.Run("update product", func(t *testing.T) {
		if len(products) > 0 && len(categories) > 0 {
			productToUpdate := products[0]
			category := categories[0]
			recorder := httptest.NewRecorder()
			body := []byte(fmt.Sprintf(`{
			  "name": "Octopus",
			  "price": 10,
			  "weight": 1000,
			  "description": "Good for health",
			  "imageUrl" : null,
			  "categoryId": "%s"
			}`, category.ID))
			request, _ := http.NewRequest("PUT", "/product/"+productToUpdate.ID, bytes.NewBuffer(body))
			request.Header.Add("Authorization", authorization())
			router.ServeHTTP(recorder, request)
			assert.Equal(t, 200, recorder.Code)
		}
	})
}
