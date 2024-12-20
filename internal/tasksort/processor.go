package tasksort

type TaskHandler struct {
	Formatter Formater
}

func (handler TaskHandler) ProcessTasks(tasks []Task) string {
	sortedTasks := Sort(tasks)
	return handler.Formatter.Format(sortedTasks)
}
