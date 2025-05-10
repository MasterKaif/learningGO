package main

import "slices"

var tasks = []Task{}
var nextID = 1

func addTask(title string) Task {
	task := Task{ID: nextID, Title: title}
	tasks = append(tasks, task)
	nextID++
	return task
}

func getTasks() []Task {
	return tasks
}

func completeTask(id int) bool {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Complete = true
			return true
		}
	}
	return false
}

func deleteTask(id int) bool {
	for i := range tasks {
		if tasks[i].ID == id {
			// tasks = append(tasks[:i], tasks[i+1:]...)
			tasks = slices.Delete(tasks, i, i+1)
			return true
		}
	}
	return false
}

