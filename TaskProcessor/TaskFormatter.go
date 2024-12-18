package TaskProcessor

import (
	"encoding/json"
	"net/http"
	"task-handler/Models"
)

type TaskFormatter struct {
	WebContext http.ResponseWriter
}

func (formatter TaskFormatter) Format(tasks []Models.Task) {
	var formattedTasks []Models.TaskJSONResponse
	for _, task := range tasks {
		outputTask := Models.TaskJSONResponse {Name: task.Name, Command: task.Command}
		formattedTasks = append(formattedTasks, outputTask)
	}
	formatter.WebContext.Header().Set("Content-Type", "application/json")
	json.NewEncoder(formatter.WebContext).Encode(formattedTasks)
}
