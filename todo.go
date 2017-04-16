package main

import "time"

// Todo struct definition
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos type definition
type Todos []Todo
