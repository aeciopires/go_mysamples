package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type messageData struct {
	Message string `json:"message"`
}

func postDataAndReturnResponse(site string, msg messageData) messageData {
	jsonBytes, _ := json.Marshal(msg)
	r, err := http.Post(site, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	message := messageData{}
	err = json.Unmarshal(data, &message)
	if err != nil {
		log.Fatal(err)
	}

	return message
}

func main() {
	siteURL := "http://localhost:8080"
	msg := messageData{Message: "Hi Server!"}
	data := postDataAndReturnResponse(siteURL, msg)
	fmt.Println(data.Message)
}
