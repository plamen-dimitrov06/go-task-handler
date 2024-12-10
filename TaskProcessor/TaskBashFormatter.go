package TaskProcessor

import (
	"fmt"
	"strings"
	"net/http"
	"task-handler/Models"
)

type TaskBashFormatter struct {
	WebContext http.ResponseWriter
}

func (formatter TaskBashFormatter) Format(tasks []Models.Task) {
	var formattedTasks []string
	formattedTasks = append(formattedTasks, "#!/usr/bin/env bash\n")
	for _, task := range tasks {
		formattedTasks = append(formattedTasks, task.Command)
	}
	formatter.WebContext.Header().Set("Content-Type", "text/html")
	taskList := strings.Join(formattedTasks, "\n")
	fmt.Fprintf(formatter.WebContext, "%v", taskList)
}