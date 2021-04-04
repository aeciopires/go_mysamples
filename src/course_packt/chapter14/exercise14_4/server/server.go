package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
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

func main() {
	listen := ":8080"
	log.Fatal(http.ListenAndServe(listen, server{}))
}
