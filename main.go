package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/todos", toDoIndex)
	router.HandleFunc("/todos/{todoId}", toDoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func toDoIndex(w http.ResponseWriter, r *http.Request) {
	todos := ToDos{
		ToDo{Name: "Write presentation"},
		ToDo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}

func toDoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	toDoID := vars["toDoID"]
	fmt.Fprintln(w, "ToDo show:", toDoID)
}
