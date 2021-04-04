package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Names struct {
	Names []string `json:"names"`
}

func getDataAndParseResponse(site string) (int, int) {
	r, err := http.Get(site)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	names := Names{}
	err = json.Unmarshal(data, &names)
	if err != nil {
		log.Fatal(err)
	}

	electricCount := 0
	boogalooCount := 0

	for _, name := range names.Names {
		if name == "Electric" {
			electricCount++
		} else if name == "Boogaloo" {
			boogalooCount++
		}
	}

	return electricCount, boogalooCount
}

func main() {
	siteURL := "https://www.google.com"

	electricCount, boogalooCount := getDataAndParseResponse(siteURL)
	fmt.Println("Electric Count: ", electricCount)
	fmt.Println("Boogaloo Count: ", boogalooCount)
}
