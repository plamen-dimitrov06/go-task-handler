package TaskProcessor

import (
	"task-handler/Contracts"
	"task-handler/Models"
)

type TaskHandler struct {
	Sorter Contracts.Sorter
	Formatter Contracts.Formater
}

func (handler TaskHandler) ProcessTasks(tasks []Models.Task) {
	sortedTasks := handler.Sorter.Sort(tasks)
	handler.Formatter.Format(sortedTasks)
}
