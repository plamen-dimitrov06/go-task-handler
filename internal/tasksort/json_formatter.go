package tasksort

import (
	"encoding/json"
	"net/http"
)

type JSONFormatter struct {
}

func (formatter JSONFormatter) Format(tasks []Task) []JSONResponse {
	formattedTasks := []JSONResponse
	for _, task := range tasks {
		outputTask := JSONResponse{Name: task.Name, Command: task.Command}
		formattedTasks = append(formattedTasks, outputTask)
	}
	return formattedTasks
}
