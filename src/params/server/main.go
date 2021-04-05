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
	cd server
	go mod init server
	go get github.com/spf13/viper
*/

package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	formName := "myFile"
	var permission fs.FileMode = 0764

	uploadedFile, uploadedFileHeader, err := r.FormFile(formName)
	if err != nil {
		log.Fatal(err)
	}

	defer uploadedFile.Close()
	fileContent, err := ioutil.ReadAll(uploadedFile)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("./%s", uploadedFileHeader.Filename), fileContent, permission)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("%s Uploaded!", uploadedFileHeader.Filename)))
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

	listen := getVariableString("SERVER_PORT", "server.port", "localhost:9000")

	fmt.Println("[DEBUG] server.port:\t", listen)
	//------------- END-VARIABLES -------------

	log.Fatal(http.ListenAndServe(listen, server{}))
}
