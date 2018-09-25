package main

import (
	"fmt"
	"time"
)

type todo struct {
	id          int
	title       string
	description string
	priority    uint8
	due         time.Time
	estlength   time.Duration
	timetaken   time.Duration
	snoozed     bool
	created     time.Time
	complete    bool
}

func main() {
	todo1 := todo{
		id:          0,
		title:       "Finish writing todo structs in go",
		description: "",
		priority:    1,
		created:     time.Now(),
		complete:    false,
		estlength:   60000000000,
	}
	fmt.Println(todo1.title)
}
