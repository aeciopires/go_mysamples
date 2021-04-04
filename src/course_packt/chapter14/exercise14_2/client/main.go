package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Type messageData in JSON format
type messageData struct {
	Message string `json:"message"`
}

// Func GetDataAndReturnResponse receive one site URL and send content of page
func GetDataAndReturnResponse(site string) messageData {
	response, err := http.Get(site)
	if err != nil {
		log.Fatal(err)
	}

	// Close for last the read of content of Body
	defer response.Body.Close()
	// Get data from the response body
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Filter JSON content for get message
	message := messageData{}
	err = json.Unmarshal(data, &message)
	if err != nil {
		log.Fatal(err)
	}

	return message
}

func main() {
	siteURL := "http://localhost:8080"
	data := GetDataAndReturnResponse(siteURL)
	fmt.Println(data.Message)
}
