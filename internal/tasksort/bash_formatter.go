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
	formattedTasks := []string
	formattedTasks = append(formattedTasks, "#!/usr/bin/env bash\n")
	for _, task := range tasks {
		formattedTasks = append(formattedTasks, task.Command)
	}
	return taskList := strings.Join(formattedTasks, "\n")
}
