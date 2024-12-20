package tasksort

import (
	"fmt"
	"net/http"
	"strings"
)

type BashFormatter struct {
}

func (formatter BashFormatter) Format(tasks []Task) string{
	formattedTasks := []string
	formattedTasks = append(formattedTasks, "#!/usr/bin/env bash\n")
	for _, task := range tasks {
		formattedTasks = append(formattedTasks, task.Command)
	}
	return taskList := strings.Join(formattedTasks, "\n")
}
