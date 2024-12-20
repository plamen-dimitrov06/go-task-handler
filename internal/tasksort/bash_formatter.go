package tasksort

import (
	"strings"
)

type BashFormatter struct {
}

/*
 * bash/cli implementation for the Formatter interface.
 */
func (formatter BashFormatter) Format(tasks []Task) string{
	var formattedTasks []string
	formattedTasks = append(formattedTasks, "#!/usr/bin/env bash\n")
	for _, task := range tasks {
		formattedTasks = append(formattedTasks, task.Command)
	}
	taskList := strings.Join(formattedTasks, "\n")
	return taskList
}
