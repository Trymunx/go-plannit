package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	todo1 := &Todo{
		id:          0,
		title:       "Finish writing todo structs in go",
		description: "",
		priority:    1,
		created:     time.Now(),
		complete:    false,
		estlength:   60000000000,
	}
	fmt.Println(todo1.title)

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/todos/:todo", GetTodos)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}
