package TaskProcessor

import (
	"encoding/json"
	"net/http"
	"fmt"
	"strings"
)

type Task struct {
	Name string `json:"name"`
	Command string `json:"command"`
	Requires []string `json:"requires"`
}

type TaskRequest struct {
	Tasks []Task `json:"tasks"`
}

type TaskOutputFormat struct {
	Name string `json:"name"`
	Command string `json:"command"`
}

func SortTasks(w http.ResponseWriter, r *http.Request) {
	var taskReq TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&taskReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

    var sortedTasks = SortTasksInternally(taskReq.Tasks)
	var formattedTasks = FormatTasks(sortedTasks)
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(formattedTasks)
}

func SortTasksInternally(tasks []Task) []Task {
	for index, task := range tasks {
		if len(task.Requires) > 0 {
			for _, requiredTask := range task.Requires {
				for targetTaskIndex, targetTask := range tasks {
					if n := strings.Compare(requiredTask, targetTask.Name); n == 0 && index < targetTaskIndex {
						fmt.Printf("Swapping %s with %s", task.Name, targetTask.Name)
						fmt.Println()
						tasks[index], tasks[targetTaskIndex] = tasks[targetTaskIndex], tasks[index]
					}
				}
			}
		}
	}
	return tasks
}

func FormatTasks(tasks []Task) []TaskOutputFormat {
	var formattedTasks []TaskOutputFormat
	for _, task := range tasks {
		outputTask := TaskOutputFormat {Name: task.Name, Command: task.Command}
		formattedTasks = append(formattedTasks, outputTask)
	}
	return formattedTasks
}