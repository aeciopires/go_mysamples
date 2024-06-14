/*
Author: Aécio Pires
Date: 2024-06-14

Description:
--------------------------------
	This program generates an HTML file listing the files in a directory and serves it over HTTP.


Usage:
--------------------------------
	To run the program, use the following command:

	go run generateIndexFiles.go <directory>

	Replace <directory> with the path to the directory you want to list files from.

	For example:

	go run generateIndexFiles.go

	This will start an HTTP server on port 8080 that lists the files in the specified directory.
	Open a web browser and navigate to http://localhost:8080/ to see the list of files.
	The HTML structure for listing files can be customized by editing the index.tmpl file.

Credit:
--------------------------------
	Generated using Google Gemini, ChatGPT and GitHub Copilot

External packages:
--------------------------------
	go get github.com/spf13/viper@v1.19.0
	go get github.com/spf13/pflag

Requirements:
--------------------------------
	- Go 1.22 or later

Creating the index.tmpl file:
--------------------------------
	Create the index.tmpl file with your desired HTML structure for listing files.
	You can use basic HTML with unordered lists (<ul>) and list items (<li>).
	For example:

```html
<!DOCTYPE html>
<html lang="en">
<head>
		<meta charset="UTF-8">
		<title>Files in {{.Dir}}</title>
</head>
<body>
		<h1>Files in {{.Dir}}</h1>
		<ul>
				{{range $index, $link := .Links}}
				<li><a href="{{$link}}">{{index $.Filenames $index}}</a></li>
				{{end}}
		</ul>
</body>
</html>
```

	This template assumes you have a placeholder `{{.Dir}}` for the directory path
	and a loop `{{range .Files}}` to iterate over the `Files` list containing file paths.

Creating the config.yaml file:
--------------------------------
	This program use ``viper`` package to read configuration values in the following order:
		* command-line arguments;
		* environment variables (INDEX_PORT and INDEX_DIRECTORY);
		* configuration file (config.yaml);
		* default values.
	The configuration file should be named config.yaml and placed in the same directory as the executable.
	Default values are defined for the port and directory.
	The handler function retrieves the directory path from the configuration.
	The main function reads the configuration and starts the server on the configured port.
	Here's an example config.yaml file:

```yaml
port: "8080"
directory: "/path/to/your/directory"
```
You can adjust the values in config.yaml as needed. The server will now use the configured HTTP port and directory.
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
	"regexp"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// FileData holds information about a file.
type FileData struct {
	Path string
}

// extractFilename extracts the filename from a URL link.
func extractFilename(link string) string {
	// Define a regular expression to match the filename portion of the link
	re := regexp.MustCompile(`[^/]+$`)
	// Find the first match of the regex in the link
	match := re.FindStringSubmatch(link)
	// If there's a match, return the captured group (filename)
	if match != nil {
		return match[0]
	}
	// Otherwise, return the original link
	return link
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

	// Define a FuncMap with the extractFilename function.
	funcMap := template.FuncMap{
		"extractFilename": extractFilename,
	}

	// Parse the template with the FuncMap.
	return template.New("index").Funcs(funcMap).Parse(string(data))
}

// handler is an HTTP handler that lists files in the given directory.
func handler(w http.ResponseWriter, r *http.Request) {
	// If the URL path is not "/", return a 404 Not Found error.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Get the directory path from the configuration.
	dir := viper.GetString("directory")
	// List the files in the directory.
	files, err := listFiles(dir)
	// If there is an error, log it and return a 500 Internal Server Error.
	if err != nil {
		log.Println(err)
		http.Error(w, "Error listing files", http.StatusInternalServerError)
		return
	}

	// Create a list of links to download the files.
	var links []string
	var filenames []string
	// Iterate over the files and create a link for each file.
	for _, file := range files {
		link := fmt.Sprintf("/download-symlink?path=%s", file.Path) // Create link to download script
		links = append(links, link)
		filenames = append(filenames, extractFilename(file.Path))
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
		Links     []string // Add links to download files
		Filenames []string // Add filenames for display
		Dir       string   // Add directory path as a field
	}{links, filenames, dir})
	// If there is an error, log it and return a 500 Internal Server Error.
	if err != nil {
		log.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

// downloadHandler serves the requested files.
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// Get the file path from the URL query parameters.
	filePath := r.URL.Query().Get("path")
	if filePath == "" {
		http.Error(w, "File path is missing", http.StatusBadRequest)
		return
	}

	// Open the file.
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Get the file's information.
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Could not get file info", http.StatusInternalServerError)
		return
	}

	// Set the appropriate headers.
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileInfo.Name()))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	// Serve the file.
	http.ServeFile(w, r, filePath)
}

// showHelp displays the help message.
func showHelp() {
	fmt.Println("Usage: go run generateIndexFiles.go [flags]")
	fmt.Println("Flags:")
	pflag.PrintDefaults()
}

// main is the entry point for the program.
func main() {
	// Define command-line flags.
	pflag.StringP("port", "p", "8080", "Port for the HTTP server")
	pflag.StringP("directory", "d", ".", "Directory to list files from")
	configFile := pflag.StringP("config", "c", "", "Path to the config file")
	help := pflag.BoolP("help", "h", false, "Show help message")
	pflag.Parse()

	// Show help message if --help or -h is provided.
	if *help {
		showHelp()
		return
	}

	// Bind flags to viper.
	viper.BindPFlag("port", pflag.Lookup("port"))
	viper.BindPFlag("directory", pflag.Lookup("directory"))

	// Bind environment variables to keys
	viper.BindEnv("port", "INDEX_PORT")
	viper.BindEnv("directory", "INDEX_DIRECTORY")

	// Set the config file if provided.
	if *configFile != "" {
		viper.SetConfigFile(*configFile)
	} else {
		// Set the name and type of the config file.
		viper.SetConfigFile("config.yaml") // name of config file (with extension)
		//viper.SetConfigName("config") // name of config file (without extension)
		//viper.SetConfigType("yaml")   // required if the config file does not have the extension in the name
		viper.AddConfigPath(".") // optionally look for config in the working directory
	}

	viper.AutomaticEnv() // read in environment variables that match

	// Default values
	viper.SetDefault("port", "8080")
	viper.SetDefault("directory", ".")

	// Read the config file if it exists.
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Printf("Trying %s. Using default values", err)
	}

	// Use the configuration values.
	port := viper.GetString("port")
	directory := viper.GetString("directory")

	// The handler function reads the index.tmpl file and renders it with the list of files in the given directory.
	http.HandleFunc("/", handler)
	// The downloadHandler function serves the requested files.
	http.HandleFunc("/download-symlink", downloadHandler)
	// Start the server on the configured port.
	fmt.Printf("Server listening on port %s, serving directory %s\n", port, directory)
	// The server will list files in the directory specified in the configuration.
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
