/*
Author: Aécio Pires
Date: 2021-09-06

Description:
	This program generates an HTML file listing the files in a directory and serves it over HTTP.

Summary:
	The program reads the directory path from the command-line arguments and lists the files in that directory.

Usage:
	To run the program, use the following command:

	go run generateIndexFiles.go <directory>

	Replace <directory> with the path to the directory you want to list files from.

	For example:

	go run generateIndexFiles.go /path/to/directory

	This will start an HTTP server on port 8080 that lists the files in the specified directory.
	Open a web browser and navigate to http://localhost:8080/ to see the list of files.
	The HTML structure for listing files can be customized by editing the index.tmpl file.
	The template uses the placeholder `{{.Dir}}` for the directory path and `{{range .Files}}` to iterate over the list of files.
	You can modify the template to change the appearance of the file listing.

Credit: generated using Google Gemini and GitHub Copilot

Requirements:
	- Go 1.22 or later
*/

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type FileData struct {
	Path string
}

// listFiles returns a list of files in the given directory.
func listFiles(dir string) ([]FileData, error) {
	var files []FileData
	// Walk walks the file tree rooted at dir and calls fn for each file or directory in the tree, including dir itself.
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// If there is an error, return it.
		if err != nil {
			return err
		}
		// If the path is not a directory, append it to the files list.
		if !info.IsDir() {
			files = append(files, FileData{Path: path})
		}
		return nil
	})
	// If there is an error, return it.
	if err != nil {
		return nil, err
	}
	// Return the list of files.
	return files, nil
}

// getIndexTemplate reads the index.tmpl file and returns a template object.
func getIndexTemplate() (*template.Template, error) {
	// Read the index.tmpl file.
	data, err := ioutil.ReadFile("index.tmpl")
	// If there is an error, return it.
	if err != nil {
		return nil, err
	}
	// Parse the template and return it.
	return template.New("index").Parse(string(data))
}

// handler is an HTTP handler that lists files in the given directory.
func handler(w http.ResponseWriter, r *http.Request) {
	// If the URL path is not "/", return a 404 Not Found error.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Get the directory path from the command-line arguments.
	dir := os.Args[1]
	// List the files in the directory.
	files, err := listFiles(dir)
	// If there is an error, log it and return a 500 Internal Server Error.
	if err != nil {
		log.Println(err)
		http.Error(w, "Error listing files", http.StatusInternalServerError)
		return
	}

	// Get the index template.
	tmpl, err := getIndexTemplate()
	// If there is an error, log it and return a 500 Internal Server Error.
	if err != nil {
		log.Println(err)
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Render the template with the list of files.
	err = tmpl.Execute(w, struct {
		Files []FileData
		Dir   string // Add directory path as a field
	}{files, dir})
	// If there is an error, log it and return a 500 Internal Server Error.
	if err != nil {
		log.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

// main is the entry point for the program.
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run generateIndexFiles.go <directory>")
		return
	}

	// Create the index.tmpl file with your desired HTML structure for listing files.
	// You can use basic HTML with unordered lists (<ul>) and list items (<li>).
	// For example:

	// ```html
	// <!DOCTYPE html>
	// <html lang="en">
	// <head>
	//     <meta charset="UTF-8">
	//     <title>Index</title>
	// </head>
	// <body>
	//     <h1>Files in {{.Dir}}</h1>
	//     <ul>
	//         {{range .Files}}
	//         <li>{{.Path}}</li>
	//         {{end}}
	//     </ul>
	// </body>
	// </html>
	// ```

	// This template assumes you have a placeholder `{{.Dir}}` for the directory path
	// and a loop `{{range .Files}}` to iterate over the `Files` list containing file paths.

	// The handler function reads the index.tmpl file and renders it with the list of files in the given directory.
	http.HandleFunc("/", handler)
	// Start the server on port 8080.
	fmt.Println("Server listening on port 8080")
	// The server will list files in the directory specified as the first command-line argument.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
