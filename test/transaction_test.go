package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/assert/v2"
	"github.com/ramdanariadi/dot-test/transaction"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var transactions []*transaction.Transaction

func TestTransaction(t *testing.T) {
	t.Run("create transaction", func(t *testing.T) {
		if len(products) > 0 {
			recorder := httptest.NewRecorder()
			p := products[0]
			body := []byte(fmt.Sprintf(`{
				  "detailTransaction": [
					{
					  "productId": "%s",
					  "total": 2
					}
				  ]
				}`, p.ID))
			request, _ := http.NewRequest("POST", "/transaction/", bytes.NewBuffer(body))
			request.Header.Add("Authorization", authorization())
			router.ServeHTTP(recorder, request)
			assert.Equal(t, 200, recorder.Code)
		}
	})

	t.Run("get all transaction", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/transaction/", nil)
		request.Header.Add("Authorization", authorization())
		router.ServeHTTP(recorder, request)

		bytes, _ := io.ReadAll(recorder.Body)
		err := json.Unmarshal(bytes, &transactions)
		assert.Equal(t, err, nil)
		assert.Equal(t, 200, recorder.Code)
	})

	t.Run("get transaction by id", func(t *testing.T) {
		if len(transactions) > 0 {
			transactionToGet := transactions[0]
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/transaction/"+transactionToGet.ID, nil)
			request.Header.Add("Authorization", authorization())
			router.ServeHTTP(recorder, request)
			assert.Equal(t, 200, recorder.Code)
		}
	})
}
