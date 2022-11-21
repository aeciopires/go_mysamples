package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
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

var returnCodeError bool = false
var cssDir string = "html/css"
var htmlFile string = "html/index.html"

type Context struct {
	Title    string
	Hostname string
}

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

func hostname(w http.ResponseWriter, req *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	if returnCodeError {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Loading the HTML template
	parsedTemplate := template.Must(template.ParseFiles(htmlFile))

	context := Context{
		Title:    "kube-pires",
		Hostname: hostname,
	}

	errTemplate := parsedTemplate.Execute(w, context)

	if errTemplate != nil {
		log.Println("Error executing template :", errTemplate)
		return
	}
}

/*
  // Code for get interface name and IPAddress
	list, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for i, iface := range list {
		fmt.Printf("%d ifaceName=%s \n", i, iface.Name)

		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}

		for j, addr := range addrs {
			fmt.Printf(" %d %v\n", j, addr)
		}
	}
*/

func main() {
	// Variables
	listen := "0.0.0.0"
	port := "3000"
	address := listen + ":" + port

	// HTTP Handlers
	fs := http.FileServer(http.Dir(cssDir))
	http.Handle("/css/", http.StripPrefix("/css", fs))
	http.HandleFunc("/", hostname)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/health", status)

	// Init webserver
	errServer := http.ListenAndServe(address, nil)
	if errServer != nil {
		log.Fatal("[ERROR] Starting the HTTP Server :", errServer)
		return
	}
}
