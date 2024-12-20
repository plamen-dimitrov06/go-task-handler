package TaskProcessor

import (
	"encoding/json"
	"net/http"
)

type TaskFormatter struct {
	WebContext http.ResponseWriter
}

func (formatter TaskFormatter) Format(tasks []Task) {
	// @TODO dont use var in func ctx
	var formattedTasks []TaskJSONResponse
	for _, task := range tasks {
		outputTask := TaskJSONResponse{Name: task.Name, Command: task.Command}
		formattedTasks = append(formattedTasks, outputTask)
	}
	formatter.WebContext.Header().Set("Content-Type", "application/json")
	json.NewEncoder(formatter.WebContext).Encode(formattedTasks)
}
