package TaskProcessor

import (
	"task-handler/Models"
)

type TaskFormatter struct {
}

func (formatter TaskFormatter) Format(tasks []Models.Task) []Models.TaskOutputFormat {
	var formattedTasks []Models.TaskOutputFormat
	for _, task := range tasks {
		outputTask := Models.TaskOutputFormat {Name: task.Name, Command: task.Command}
		formattedTasks = append(formattedTasks, outputTask)
	}
	return formattedTasks
}