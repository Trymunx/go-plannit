package main

import (
	"fmt"
	"html"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index is the root of the file
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))
}

// GetTodos should return a list of titles
func GetTodos(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Todo: %s", ps.ByName("name"))
}

// Hello is a greety route thing
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello %s!\n", ps.ByName("name"))
}
