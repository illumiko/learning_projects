package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const PORT = "8080"

type Id string
type BookInfo struct {
	Name   string `json:"name"`
	Rating int    `json:"rating"`
	Author Author `json:"Author"`
}

type Author struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
}

type Booklist map[Id]BookInfo

var books = Booklist{
	"111": {
		Name:   "Book Of First",
		Rating: 7,
		Author: Author{"John", "Doe"},
	},
	"222": {
		Name:   "Book Of Second",
		Rating: 9,
		Author: Author{"Jake", "Doe"},
	},
}

func main() {

	http.HandleFunc("/t", GetBooks)

	fmt.Println("Starting Server @port:", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)

}
