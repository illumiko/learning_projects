package main

import (
	"encoding/json"
	"testing"
)

func TestDB(t *testing.T) {
	t.Run("teesting json db get", func(t *testing.T) {
		f := CreateDB(filename_db, `[
            {"Name":"John","Wins":"23"}
            ]`)
		f.Seek(0, 0)
		defer f.Close()

		var store db

		err := json.NewDecoder(f).Decode(&store)
		if err != nil {
			t.Fatal(err)
		}

	})
}
