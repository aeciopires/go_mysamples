package main

import (
	"embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"text/template"
)

// References:
// https://gobyexample.com/http-server
// https://github.com/cod3rcursos/curso-go/tree/master/http
// https://github.com/weaveworks/scope/blob/v1.13.2/common/hostname/hostname.go
// https://github.com/richardpct/go-hostname/blob/master/hostname02/hostname.go
// https://adlerhsieh.com/blog/rendering-dynamic-data-in-go-http-template
// https://github.com/paulbouwer/hello-kubernetes
// https://golangdocs.com/templates-in-golang
// https://hackthedeveloper.com/golang-html-template-parsefiles-and-execute/
// https://blog.gopheracademy.com/advent-2017/using-go-templates/
// https://philipptanlak.com/mastering-html-templates-in-go-the-fundamentals/
// https://pkg.go.dev/html/template#example-Template-Parsefiles
// https://stackoverflow.com/questions/61057271/how-to-run-html-with-css-using-golang
// https://stackoverflow.com/questions/28793619/golang-what-to-use-http-servefile-or-http-fileserver/28798174#28798174
// https://www.tumblr.com/golang-examples/99458329439/get-local-ip-addresses
// https://pkg.go.dev/embed
// https://charly3pins.dev/blog/learn-how-to-use-the-embed-package-in-go-by-building-a-web-page-easily/

const (
	htmlExtension = "templates/*.html"
)

var (
	// Embed static content for using with web server.
	//
	//go:embed templates/*
	filesEmbed embed.FS

	returnCodeError bool   = false
	cssDir          string = "css"
)

type (
	Content struct {
		Title     string
		Hostname  string
		Addresses map[string]string
	}
)

// Show status of application
func status(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "status: OK\n")
}

// Receive and return the headers of the HTTP requisition
func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, value := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, value)
		}
	}
}

// Get hostname
func hostname() (name string) {

	name, err := os.Hostname()

	if err != nil {
		log.Fatal(err)
	}

	return
}

// Get interface name and IPAddress
func getAddress() (address map[string]string) {

	address = make(map[string]string)

	listIfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, iface := range listIfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			log.Fatal(err)
		}

		if addrs != nil {
			address[iface.Name] = addrs[0].String()
		}
	}

	return
}

// Parse of HTML content
func parseTemplate(w http.ResponseWriter, req *http.Request) {

	if returnCodeError {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Loading the HTML template
	parsedTemplate := template.Must(
		template.New("htmlFiles").
			ParseFS(filesEmbed, "templates/*"),
	)

	content := Content{
		Title:     "kube-pires",
		Hostname:  hostname(),
		Addresses: getAddress(),
	}

	errTemplate := parsedTemplate.ExecuteTemplate(w, "index.html", content)

	if errTemplate != nil {
		log.Println("[ERROR] Executing template :", errTemplate)
		return
	}
}

func main() {

	// Variables
	listen := "0.0.0.0"
	port := "3000"
	address := listen + ":" + port

	// HTTP Handlers
	// Exposing CSS files without embed directive
	fs := http.FileServer(http.Dir(cssDir))
	http.Handle("/css/", http.StripPrefix("/css", fs))
	http.HandleFunc("/", parseTemplate)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/health", status)

	// Init webserver
	errServer := http.ListenAndServe(address, nil)
	if errServer != nil {
		log.Fatal("[ERROR] Starting the HTTP Server :", errServer)
		return
	}
}
