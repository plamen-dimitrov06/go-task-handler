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

		JSONFormater := tasksort.JSONFormater{}
		processor := tasksort.TaskHandler{Formater: JSONFormater}
		formattedTasks := processor.ProcessTasks(taskReq.Tasks)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%v", formattedTasks)
	})

	http.HandleFunc("/bash", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&taskReq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		bashFormater := tasksort.BashFormater{}
		processor := tasksort.TaskHandler{Formater: bashFormater}
		formattedTasks := processor.ProcessTasks(taskReq.Tasks)

		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "%v", formattedTasks)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
