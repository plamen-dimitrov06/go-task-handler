package tasksort

import (
	"encoding/json"
	"net/http"
)

type JSONFormatter struct {
	WebContext http.ResponseWriter
}

func (formatter JSONFormatter) Format(tasks []Task) {
	formattedTasks := []JSONResponse
	for _, task := range tasks {
		outputTask := JSONResponse{Name: task.Name, Command: task.Command}
		formattedTasks = append(formattedTasks, outputTask)
	}
	formatter.WebContext.Header().Set("Content-Type", "application/json")
	json.NewEncoder(formatter.WebContext).Encode(formattedTasks)
}
