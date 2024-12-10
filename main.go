package main

import (
    "log"
    "net/http"
    "encoding/json"
    "task-handler/TaskProcessor"
    "task-handler/Models"
)

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        var taskReq Models.TaskRequest
        if err := json.NewDecoder(r.Body).Decode(&taskReq); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        taskSorter := TaskProcessor.TaskSorter{Tasks: taskReq.Tasks}
        taskFormatter := TaskProcessor.TaskFormatter{}
        processedTasks := TaskProcessor.ProcessTasks(taskSorter, taskFormatter)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(processedTasks)
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}