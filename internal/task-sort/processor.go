package TaskProcessor

type TaskHandler struct {
	Formatter Formater
}

func (handler TaskHandler) ProcessTasks(tasks []Task) {
	sortedTasks := Sort(tasks)
	handler.Formatter.Format(sortedTasks)
}
