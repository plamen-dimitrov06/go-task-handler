package tasksort

import (
	"encoding/json"
	"net/http"
)

type JSONFormatter struct {
}

/*
 * JSON implementation for the Formatter interface.
 */
func (formatter JSONFormatter) Format(tasks []Task) []JSONResponse {
	formattedTasks := []JSONResponse
	for _, task := range tasks {
		outputTask := JSONResponse{Name: task.Name, Command: task.Command}
		formattedTasks = append(formattedTasks, outputTask)
	}
	return formattedTasks
}
