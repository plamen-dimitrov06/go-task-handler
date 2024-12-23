package tasksort

type TaskHandler struct {
	formater Formater
}

func (handler TaskHandler) ProcessTasks(tasks []Task) string {
	sortedTasks := Sort(tasks)
	return handler.formater.Format(sortedTasks)
}

func NewJSONProcessor() TaskHandler {
	JSONFormater := NewJSONFormater()
	return TaskHandler{formater: JSONFormater}
}

func NewBashProcessor() TaskHandler {
	bashFormater := NewBashFormater()
	return TaskHandler{formater: bashFormater}
}
