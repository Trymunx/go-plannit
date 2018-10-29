package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Route is an HTTP route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is the list of all routes
type Routes []Route

// NewRouter is how we create a router struct
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return nil
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetTodos",
		"GET",
		"/todos/:todo",
		GetTodos,
	},
	Route{
		"Greet",
		"GET",
		"/hello/:name",
		Hello,
	},
}
