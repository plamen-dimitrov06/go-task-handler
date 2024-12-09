package TaskProcessor

type TaskFormatter struct {
	Tasks []Task
}

func (formater TaskFormatter) format() []TaskOutputFormat {
	var formattedTasks []TaskOutputFormat
	for _, task := range formater.Tasks {
		outputTask := TaskOutputFormat {Name: task.Name, Command: task.Command}
		formattedTasks = append(formattedTasks, outputTask)
	}
	return formattedTasks
}