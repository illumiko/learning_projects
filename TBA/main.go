package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	iData := `[
    {"Name":"John","Wins":"23"},
    {"Name":"Joe","Wins":"21"}
    ]`
	f := CreateDB(filename_db, iData)
	defer f.Close()

	store := dbStore{}

	f.Seek(0, 0)
	var decoded db
	err := json.NewDecoder(f).Decode(&decoded)
	if err != nil {
		log.Fatal("json read err", err)
	}
	fmt.Println(decoded.Get("Joe"))

}
