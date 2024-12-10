package TaskProcessor

import (
	"task-handler/Contracts"
)

func ProcessTasks(sorter Contracts.Sorter, formatter Contracts.Formater) {
	sortedTasks := sorter.Sort()
	formatter.Format(sortedTasks)
}