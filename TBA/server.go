package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Store interface {
	Get(id string) int
	Update(name string, value string)
}

type Server struct {
	store Store
	http.Handler
}

func NewServer(store Store) *Server {
	p := new(Server)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/get/", http.HandlerFunc(p.handleGet))
	router.Handle("/update/", http.HandlerFunc(p.handleUpdate))

	p.Handler = router
	return p

}

func (s *Server) handleGet(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/get/")
	w.WriteHeader(http.StatusOK)
	value := s.store.Get(id)
	json.NewEncoder(w).Encode(value)

	// fmt.Fprint(w, value)
}

func (s *Server) handleUpdate(w http.ResponseWriter, r *http.Request) {
	parsedURI, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Fatalf("Error parsing URL: %v", err)
	}

	w.WriteHeader(http.StatusAccepted)
	id := parsedURI.Get("name")
	value := parsedURI.Get("value")

	s.store.Update(id, value)
	// x, _ := url.ParseQuery(id.RawQuery)
	// value := strings.TrimPrefix(id, , id))
}
