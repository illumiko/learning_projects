package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	text := "hello"
	w.Write([]byte(text))
}

func main() {
	s := http.Server{Addr: ":8080", Handler: http.HandlerFunc(handler)}
	log.Fatal(s.ListenAndServe())

	fmt.Println("helo")
}
