/*
	References:
	* https://medium.com/@bnprashanth256/reading-configuration-files-and-environment-variables-in-go-golang-c2607f912b63
	* https://github.com/BNPrashanth/go-poc-bp/tree/env-var-approach1/src
	* https://schadokar.dev/posts/use-environment-variable-in-your-next-golang-project/
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

func getVariableString(envConfig string, paramConfig string, defaultValue string) string {
	if viper.GetString(envConfig) == "" {
		if viper.GetString(paramConfig) == "" {
			// Set undefined variables
			return defaultValue
		} else {
			return viper.GetString(paramConfig)
		}
	} else {
		return viper.GetString(envConfig)
	}
}

func getVariableBool(envConfig string, paramConfig string, defaultValue bool) bool {
	aux := viper.GetBool(envConfig)
	if &aux == nil {
		aux = viper.GetBool(paramConfig)
		if &aux == nil {
			// Set undefined variables
			return defaultValue
		} else {
			return viper.GetBool(paramConfig)
		}
	} else {
		return viper.GetBool(envConfig)
	}
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

	// Find and read the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("[ERROR] Error while reading config file %s", err)
	}

	protocol := getVariableString("SITE_PROTOCOL", "site.protocol", "http")
	siteAddress := getVariableString("SITE_ADDRESS", "site.address", "localhost:9000")
	myFile := getVariableString("UPLOAD_FILE", "upload.file", "test3.txt")
	siteURL := protocol + "://" + siteAddress

	fmt.Println("[DEBUG] siteURL: ", siteURL)
	fmt.Println("[DEBUG] myFile:  ", myFile)
	//------------- END-VARIABLES -------------

	data := postFileAndReturnResponse(siteURL, myFile)
	fmt.Println(data)
}
