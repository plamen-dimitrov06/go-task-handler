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

func SortTasks(w http.ResponseWriter, r *http.Request) {
	var taskReq TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&taskReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

    SortTasksInternally(taskReq.Tasks)
}

func SortTasksInternally(tasks []Task) {
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
}