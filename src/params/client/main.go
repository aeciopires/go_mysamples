/*
	References:
	* https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66
	* https://github.com/spf13/viper

	Viper uses the following precedence order. Each item takes precedence over the item below it:

	* explicit call to Set
	* flag
	* env
	* config
	* key/value store
	* default

	Install requirements:
	cd client
	go mod init client
	go get github.com/spf13/viper
*/

package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

// Configurations exported
type Configurations struct {
	Site   ServerConfigurations
	Upload UploadConfigurations
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Address  string
	Protocol string
}

// UploadConfigurations exported
type UploadConfigurations struct {
	File string
}

func postFileAndReturnResponse(site string, filename string) string {
	fileDataBuffer := bytes.Buffer{}
	multipartWritter := multipart.NewWriter(&fileDataBuffer)
	formName := "myFile"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	formFile, err := multipartWritter.CreateFormFile(formName, file.Name())
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(formFile, file)
	if err != nil {
		log.Fatal(err)
	}
	multipartWritter.Close()

	req, err := http.NewRequest("POST", site, &fileDataBuffer)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", multipartWritter.FormDataContentType())
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func main() {

	//------------- BEGIN-VARIABLES -------------
	// Set the file name of the configurations file
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	var configuration Configurations

	// Find and read the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("[ERROR] Error while reading config file %s", err)
	}

	err2 := viper.Unmarshal(&configuration)
	if err2 != nil {
		fmt.Printf("[ERROR] Unable to decode into struct, %v", err2)
	}

	protocol := viper.GetString("site.protocol")
	if protocol == "" {
		protocol = viper.GetString("SITE_PROTOCOL")
		if protocol == "" {
			// Set undefined variables
			protocol = "http"
		}
	}

	siteAddress := viper.GetString("site.address")
	if siteAddress == "" {
		siteAddress = viper.GetString("SITE_ADDRESS")
		if siteAddress == "" {
			// Set undefined variables
			siteAddress = "localhost:9000"
		}
	}

	siteURL := protocol + "://" + siteAddress

	myFile := viper.GetString("upload.file")
	if myFile == "" {
		myFile = viper.GetString("UPLOAD_FILE")
		if myFile == "" {
			// Set undefined variables
			myFile = "test3.txt"
		}
	}

	fmt.Println("[DEBUG] siteURL:\t", siteURL)
	fmt.Println("[DEBUG] myFile:\t", myFile)
	//------------- END-VARIABLES -------------

	data := postFileAndReturnResponse(siteURL, myFile)
	fmt.Println(data)
}
