package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	taskReq := Models.TaskRequest{}
	taskSorter := TaskProcessor.TaskSorter{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&taskReq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		taskFormatter := TaskProcessor.TaskFormatter{WebContext: w}
		processor := TaskProcessor.TaskHandler{Sorter: taskSorter, Formatter: taskFormatter}
		processor.ProcessTasks(taskReq.Tasks)
	})

	http.HandleFunc("/bash", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&taskReq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		taskFormatter := TaskProcessor.TaskBashFormatter{WebContext: w}
		processor := TaskProcessor.TaskHandler{Sorter: taskSorter, Formatter: taskFormatter}
		processor.ProcessTasks(taskReq.Tasks)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
