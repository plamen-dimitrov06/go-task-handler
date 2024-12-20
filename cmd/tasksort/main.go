package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tasksort/internal/tasksort"
)

func main() {
	taskReq := tasksort.TaskRequest{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&taskReq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		taskFormatter := tasksort.JSONFormater{}
		processor := tasksort.TaskHandler{Formater: taskFormatter}
		formattedTasks := processor.ProcessTasks(taskReq.Tasks)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%v", formattedTasks)
	})

	http.HandleFunc("/bash", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&taskReq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		taskFormatter := tasksort.BashFormater{}
		processor := tasksort.TaskHandler{Formater: taskFormatter}
		formattedTasks := processor.ProcessTasks(taskReq.Tasks)

		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "%v", formattedTasks)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
