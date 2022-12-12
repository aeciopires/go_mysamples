package main

import (
	"context"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

// An empty struct postsResource is defined to namespace all functionality
// related to "posts" resources.
type postsResource struct{}

// Defines the Routes method on the struct postsResource. This method returns
// an instance of the chi router that handles the /posts route and its sub-routes.
// Recall previously how we attached this sub-router to the parent router in main.go
func (resource postsResource) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", resource.List)    // GET /posts - Read a list of posts.
	router.Post("/", resource.Create) // POST /posts - Create a new post.

	router.Route("/{id}", func(router chi.Router) {
		router.Use(PostCtx)
		router.Get("/", resource.Get)       // GET /posts/{id} - Read a single post by :id.
		router.Put("/", resource.Update)    // PUT /posts/{id} - Update a single post by :id.
		router.Delete("/", resource.Delete) // DELETE /posts/{id} - Delete a single post by :id.
	})

	return router
}

// Request Handler - GET /posts - Read a list of posts.
func (resource postsResource) List(writeResponse http.ResponseWriter, request *http.Request) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		http.Error(writeResponse, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	writeResponse.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(writeResponse, resp.Body); err != nil {
		http.Error(writeResponse, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Request Handler - POST /posts - Create a new post.
func (resource postsResource) Create(writeResponse http.ResponseWriter, request *http.Request) {
	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", request.Body)

	if err != nil {
		http.Error(writeResponse, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	writeResponse.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(writeResponse, resp.Body); err != nil {
		http.Error(writeResponse, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writeResponse http.ResponseWriter, request *http.Request) {
		ctx := context.WithValue(request.Context(), "id", chi.URLParam(request, "id"))
		next.ServeHTTP(writeResponse, request.WithContext(ctx))
	})
}

// Request Handler - GET /posts/{id} - Read a single post by :id.
func (resource postsResource) Get(writeResponse http.ResponseWriter, request *http.Request) {
	id := request.Context().Value("id").(string)

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + id)

	if err != nil {
		http.Error(writeResponse, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	writeResponse.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(writeResponse, resp.Body); err != nil {
		http.Error(writeResponse, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Request Handler - PUT /posts/{id} - Update a single post by :id.
func (resource postsResource) Update(writeResponse http.ResponseWriter, request *http.Request) {
	id := request.Context().Value("id").(string)
	client := &http.Client{}

	req, err := http.NewRequest("PUT", "https://jsonplaceholder.typicode.com/posts/"+id, request.Body)
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		http.Error(writeResponse, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		http.Error(writeResponse, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	writeResponse.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(writeResponse, resp.Body); err != nil {
		http.Error(writeResponse, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Request Handler - DELETE /posts/{id} - Delete a single post by :id.
func (resource postsResource) Delete(writeResponse http.ResponseWriter, request *http.Request) {
	id := request.Context().Value("id").(string)
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/posts/"+id, nil)

	if err != nil {
		http.Error(writeResponse, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		http.Error(writeResponse, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	writeResponse.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(writeResponse, resp.Body); err != nil {
		http.Error(writeResponse, err.Error(), http.StatusInternalServerError)
		return
	}
}
