package tasksort

type TaskHandler struct {
	Formater Formater
}

func (handler TaskHandler) ProcessTasks(tasks []Task) string {
	sortedTasks := Sort(tasks)
	return handler.Formater.Format(sortedTasks)
}
