package main

import "time"

type ToDo struct {
	Name      string
	Completed bool
	Due       time.Time
}

type ToDos []ToDo
