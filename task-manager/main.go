package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Done   bool   `json:"done"`
}

var tasks = []Task{}
var nextID = 1

func main() {
	http.HandleFunc("/tasks", handleTasks)
	http.HandleFunc("/tasks/", handleTaskByID)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(tasks)
	case "POST":
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		task.ID = nextID
		nextID++
		tasks = append(tasks, task)
		json.NewEncoder(w).Encode(task)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleTaskByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	for i := range tasks {
		if tasks[i].ID == id {
			fmt.Println("Task found:", tasks[i])
			switch r.Method {
			case "GET":
				json.NewEncoder(w).Encode(tasks[i])
				return
			case "PUT":
				var updated Task
				if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
					http.Error(w, "Invalid input", http.StatusBadRequest)
					return
				}
				updated.ID = id
				tasks[i] = updated
				json.NewEncoder(w).Encode(updated)
				return
			case "DELETE":
				tasks = append(tasks[:i], tasks[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}
