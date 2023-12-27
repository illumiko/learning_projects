package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

const (
	fname       = "./data.json"
	url_base    = "https://xkcd.com/"
	url_suffix  = "/info.0.json"
	no_of_comic = 2400
)

type comic_json struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	Safe_title string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

var comic_json_list = []comic_json{}

func main() {

	if _, err := os.Stat(fname); errors.Is(err, os.ErrNotExist) {
		var wg sync.WaitGroup

		wg.Add(no_of_comic)
		for i := 0; i < no_of_comic; i++ {
			go func(i int) {
				url := fmt.Sprint(url_base, i, url_suffix)

				resp, err := http.Get(url)
				if err != nil {
					log.Fatal(err)
				}

				if resp.StatusCode > 299 {
					fmt.Printf("%v Comic not found Status %v\n", i, resp.Status)
					wg.Done()
				} else {
					data, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						log.Fatal(err)
					}
					unmarshalled_data := &comic_json{}
					err = json.Unmarshal(data, unmarshalled_data)
					if err != nil {
						log.Fatal(err)
					}
					comic_json_list = append(comic_json_list, *unmarshalled_data)
					defer wg.Done()
				}

			}(i)

		}
		wg.Wait()
		fmt.Println("***Done Downloading***")

		file, err := os.Create(fname)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		fmt.Printf("%#v", comic_json_list)

		enc := json.NewEncoder(file)
		enc.Encode(comic_json_list)

		// file.Write([]byte("["))
		// for _, v := range comic_json_list {
		// 	file.Write(v)
		// 	file.Write([]byte(","))
		// }
		// file.Write([]byte("]"))
	} else {
		fmt.Println("Data already exists at", fname)
	}

}
