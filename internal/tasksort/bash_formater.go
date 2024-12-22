package tasksort

import (
	"strings"
)

type BashFormater struct {
}

/*
 * bash/cli implementation for the Formater interface.
 */
func (formater BashFormater) Format(tasks []Task) string {
	var formatedTasks []string
	formatedTasks = append(formatedTasks, "#!/usr/bin/env bash\n")
	for _, task := range tasks {
		formatedTasks = append(formatedTasks, task.Command)
	}
	taskList := strings.Join(formatedTasks, "\n")
	return taskList
}

func NewBashFormater() BashFormater { return BashFormater{} }
