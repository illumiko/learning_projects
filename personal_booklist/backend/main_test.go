package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPI(t *testing.T) {

	t.Run("Testing GET request", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "http://example.go/foo", nil)

		GetBooks(rr, req)

		data := rr.Result().Body
		// jdec := json.NewDecoder(data).Decode(&Booklist{})
		fmt.Println(data)

		if http.StatusOK != rr.Result().StatusCode {
			t.Errorf("\nexpected status %v\ngot status %v", http.StatusOK, rr.Result().StatusCode)
		}

	})

}
