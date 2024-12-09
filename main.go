package main

import (
    "log"
    "net/http"
    "task-handler/TaskProcessor"
)

func main() {
    http.HandleFunc("/", TaskProcessor.SortTasks)
    log.Fatal(http.ListenAndServe(":8080", nil))
}