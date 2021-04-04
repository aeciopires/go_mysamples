/*
	References:
	* https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66
	* https://github.com/spf13/viper

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

// Configurations exported
type Configurations struct {
	Server               ServerConfigurations
	SERVER_PORT          string
	SERVER_HTTP_ENABLED  bool
	SERVER_HTTPS_ENABLED bool
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port string
}

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

	listen := viper.GetString("server.port")
	if listen == "" {
		listen = viper.GetString("SERVER_PORT")
		if listen == "" {
			// Set undefined variables
			listen = "localhost:9000"
		}
	}

	fmt.Println("[DEBUG] server.port:\t", listen)
	//------------- END-VARIABLES -------------

	log.Fatal(http.ListenAndServe(listen, server{}))
}
