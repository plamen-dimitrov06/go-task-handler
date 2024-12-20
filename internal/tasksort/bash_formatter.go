package tasksort

import (
	"fmt"
	"net/http"
	"strings"
)

type BashFormatter struct {
	WebContext http.ResponseWriter
}

func (formatter BashFormatter) Format(tasks []Task) {
	formattedTasks := []string
	formattedTasks = append(formattedTasks, "#!/usr/bin/env bash\n")

	for _, task := range tasks {
		formattedTasks = append(formattedTasks, task.Command)
	}
	formatter.WebContext.Header().Set("Content-Type", "text/html")
	taskList := strings.Join(formattedTasks, "\n")
	fmt.Fprintf(formatter.WebContext, "%v", taskList)
}
