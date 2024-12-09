package TaskProcessor

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type Task struct {
	Name string `json:"name"`
	Command string `json:"command"`
	Requires []string `json:"requires"`
}

type TaskBody struct {
	Tasks []Task `json:"tasks"`
}

func SortTasks(w http.ResponseWriter, r *http.Request) {
	var tasks TaskBody
	if err := json.NewDecoder(r.Body).Decode(&tasks); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

    SortTasksInternally(tasks.Tasks)
}

func SortTasksInternally(tasks []Task) {
	fmt.Println("test2")
    fmt.Printf("Tasks: %+v", tasks)
}