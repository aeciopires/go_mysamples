// Reference: https://learning.oreilly.com/library/view/the-go-workshop/9781838647940/C14177_14_Final_SZ_ePub.xhtml

package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Func GetDataAndReturnResponse similar operation: curl https://www.google.com
func GetDataAndReturnResponse(site string) string {
	// send the GET request
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

	// return the response data
	return string(data)
}

func main() {
	siteURL := "https://www.google.com"

	data := GetDataAndReturnResponse(siteURL)
	log.Println(data)
}
