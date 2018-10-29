package main

import (
	"time"
)

// Todo is out basic todo struct
type Todo struct {
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

// Todos should store all of our todos
type Todos []Todo
