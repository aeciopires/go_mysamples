package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func getDataWithCustomOptionsAndReturnResponse(site string) string {
	client := http.Client{Timeout: 11 * time.Second}
	req, err := http.NewRequest("POST", site, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "superSecretToken")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func main() {
	siteURL := "http://localhost:8080"

	data := getDataWithCustomOptionsAndReturnResponse(siteURL)
	fmt.Println(data)
}
