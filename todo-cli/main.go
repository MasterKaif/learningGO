package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo [add|list|don|delete] [task]")
		return
	}

	cmd := os.Args[1]
	tasks, _  := loadTasks()

	switch cmd {
		case "add":
			title := os.Args[2]
			task := Tasks{
				ID:	len(tasks)+1,
				Title: title,
			}
			tasks = append(tasks, task)
			saveTasks(tasks)
			fmt.Println("Task added")
		

		case "list":
			for _, t := range tasks {
				status := "❌"
				if t.Completed {
					status = "✅"
				}

				fmt.Printf("[%d] %s %s\n", t.ID, status, t.Title)
			}
		case "done":
			id, _ := strconv.Atoi(os.Args[2])
			for i := range tasks {
				if tasks[i].ID == id {
					tasks[i].Completed = true
				}
			}

			saveTasks(tasks)
			fmt.Println("Marked as Done.")
		case "delete":
			id, _ := strconv.Atoi(os.Args[2])
			for i := range tasks {
				if tasks[i].ID == id {
					tasks = append(tasks[:i], tasks[i+1:]...)
					break
				}
			}

			saveTasks(tasks)
			fmt.Println("Task Deleted")
		default:
			fmt.Println("Unknown command:", cmd)
	}

}




