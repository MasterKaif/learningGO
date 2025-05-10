package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/tasks", listTasksHandler).Methods("GET")
	r.HandleFunc("/tasks", addTaskHandler).Methods("POST")
	r.HandleFunc("/tasks", completeTaskHandler).Methods("PUT")
	r.HandleFunc("/tasks", deleteTaskHandler).Methods("DELETE")
	
	log.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

