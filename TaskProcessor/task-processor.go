package TaskProcessor

import (
	"task-handler/Contracts"
	"task-handler/Models"
)

func ProcessTasks(sorter Contracts.Sorter, formatter Contracts.Formater) []Models.TaskOutputFormat {
	sortedTasks := sorter.Sort()
	return formatter.Format(sortedTasks)
}