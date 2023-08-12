package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const filename_db = "db.json"

type dbEntries struct {
	Name string `json:"Name"`
	Wins string `json:"Wins"`
}

type dbStore struct {
	store *[]dbEntries
	file  *os.File
}

func CreateDB(filename string, initialData string) *os.File {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("cannot create file:", err.Error())
	}

	if _, err := file.Write([]byte(initialData)); err != nil {
		log.Fatal("Cannot write to", filename_db, err)
	}

	return file
}

func (d *dbStore) Get(id string) int {
	wins := 0
	for _, v := range *d.store {
		if v.Name == id {
			wins, _ = strconv.Atoi(v.Wins)
		}
	}
	return wins
}

func (d *dbStore) Update(id string, value string) {
	for _, v := range *d.store {
		if v.Name == id {
			v.Wins = value
		}
	}
	fmt.Println(d.store)

}
