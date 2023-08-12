package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestServer(t *testing.T) {
	data := map[string]int{
		"John": 20,
	}

	store := &testStore{data}
	server := NewServer(store)

	t.Run("Get value from existing field", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/get/John", nil)
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)
		var got int

		json.NewDecoder(resp.Body).Decode(&got)
		want := "20"

		assertStatus(t, resp.Code, http.StatusOK)
		assertStoreData(t, strconv.Itoa(got), want)
	})

	t.Run("Get value from non-existing field", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/get/Bil", nil)
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)
		var got int

		json.NewDecoder(resp.Body).Decode(&got)
		want := "0"

		assertStatus(t, resp.Code, http.StatusOK)
		assertStoreData(t, strconv.Itoa(got), want)
	})

	t.Run("Update value", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/update/?name=John&value=23", nil)
		resp := httptest.NewRecorder()
		server.ServeHTTP(resp, req)

		got := server.store.Get("John")
		want := "23"

		assertStoreData(t, strconv.Itoa(got), want)
		assertStatus(t, resp.Code, http.StatusAccepted)

	})
}

type testStore struct {
	store map[string]int
}

func (t *testStore) Get(id string) int {
	return t.store[id]
}

func (t *testStore) Update(name string, value string) {
	t.store[name], _ = strconv.Atoi(value)
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Status unmatched:got:%v\nwant:%v", got, want)
	}
}

func assertStoreData(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Unexpected mismatch\ngot:%v\nwant:%v", got, want)
	}

}
