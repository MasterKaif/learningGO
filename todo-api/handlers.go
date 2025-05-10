package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func listTasksHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getTasks())
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var t Task

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	newTask := addTask(t.Title)
	json.NewEncoder(w).Encode(newTask)
}

func completeTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	if completeTask(id) {
		w.WriteHeader(http.StatusOK)
		w.Write([] byte("Marked completed"))
		return
	}else {
		http.Error(w, "Task not Found", http.StatusNotFound)
	}
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if deleteTask(id) {
		w.WriteHeader(http.StatusOK)
		w.Write([] byte("Delete"))
		return
	}else {
		http.Error(w, "Task not Found", http.StatusNotFound)
		return
	}
}

