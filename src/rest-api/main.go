package main

/*

This RESTful API service will offer the following endpoints for accessing/manipulating "posts" resources:

* GET / - Verify whether or not the service is up and running ("health check"). Returns the "It works!" message
* GET /posts - Retrieve a list of posts.
* POST /posts - Creates a post.
* GET /posts/{id} - Retrieve a single post identified by its id.
* PUT /posts/{id} - Update a single post identified by its id.
* DELETE /posts/{id} - Delete a single post identified by its id.

A post represents an online published piece of writing. It consists of a title, content, an ID and an ID of the user who authored the post:

```json
{
  "title": "",
  "body": "",
  "userId": 1,
  "id": 2
}
```

These endpoints will interface directly with JSONPlaceholder, a fake API for testing with dummy data. For example, when it receives a GET /posts request from a user, our Go RESTful API will check for this route and execute the route handler mapped to it. This route handler will send a request to https://jsonplaceholder.typicode.com/posts and pipe the response of that request back to the user.

This project will be comprised of two Go files:

* main.go - Contains a parent router, which sub-routers mount to.
* posts.go - Contains a sub-router and its handlers for the /posts route and its sub-routes (i.e. /posts/{id}).

References:
* https://www.newline.co/@kchan/building-a-simple-restful-api-with-go-and-chi--5912c411
* Content type REST: https://restfulapi.net/content-negotiation/
* https://github.com/newline-sandbox/go-chi-restful-api
* https://gabrieltanner.org/blog/collecting-prometheus-metrics-in-golang/
* https://blog.pvincent.io/2017/12/prometheus-blog-series-part-1-metrics-and-labels/

*/

import (
	"log"
	"net/http"
	"os"
	"strings"

	chiprometheus "github.com/766b/chi-prometheus"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Set default values
var (
	protocol = "http"
	address  = "0.0.0.0"
	port     = "8080"
)

func main() {

	// Reading environment variables
	if protocolEnv := os.Getenv("PROTOCOL"); protocolEnv != "" {
		protocol = strings.ToLower(protocolEnv)
	}

	if addressEnv := os.Getenv("ADDRESS"); addressEnv != "" {
		address = strings.ToLower(addressEnv)
	}

	if portEnv := os.Getenv("PORT"); portEnv != "" {
		port = portEnv
	}

	// Defining listen and baseurl
	listen := address + ":" + port
	baseurl := protocol + "://" + address + ":" + port

	log.Printf("Starting up on %s", baseurl)

	// Creating a root route using go-chi/chi framework
	rootRouter := chi.NewRouter()

	/*
			chi framework offers a lot of useful, pre-defined middleware handlers via its
			optional middleware package.One such middleware handler is Logger, which prints
			on every incoming request:
		* The date and time the request was received.
		* The endpoint the request was sent to (HTTP method and route).
		* The IP address of the client sending the request.
		* The response status code.
		* The response size (in bytes).
		* The amount of time it took to process the request (in milliseconds).
	*/
	rootRouter.Use(middleware.Logger)
	metricsRouter := chiprometheus.NewMiddleware("metrics", 400)
	rootRouter.Use(metricsRouter)

	// Metrics endpoint in format supported by Prometheus
	rootRouter.Handle("/metrics", promhttp.Handler())

	rootRouter.Get("/", func(writeResponse http.ResponseWriter, request *http.Request) {
		writeResponse.Header().Set("Content-Type", "text/plain")
		writeResponse.Write([]byte("Its works!"))
	})

	rootRouter.Get("/status", func(writeResponse http.ResponseWriter, request *http.Request) {
		writeResponse.Header().Set("Content-Type", "text/plain")
		writeResponse.Write([]byte("status: OK!"))
	})

	rootRouter.Get("/notfound", func(writeResponse http.ResponseWriter, request *http.Request) {
		writeResponse.WriteHeader(http.StatusNotFound)
		writeResponse.Write([]byte("not found"))
	})

	rootRouter.Mount("/posts", postsResource{}.Routes())

	log.Fatal(http.ListenAndServe(listen, rootRouter))
}
