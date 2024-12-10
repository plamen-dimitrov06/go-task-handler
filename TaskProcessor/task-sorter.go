package TaskProcessor

import (
	"strings"
	"task-handler/Models"
)

type TaskSorter struct {
	Tasks []Models.Task
}

/**
 * Sort tasks taking into consideration the required property.
 */
func (sorter TaskSorter) Sort() []Models.Task {
	for index, task := range sorter.Tasks {
		if len(task.Requires) > 0 {
			for _, requiredTask := range task.Requires {
				for targetTaskIndex, targetTask := range sorter.Tasks {
					if n := strings.Compare(requiredTask, targetTask.Name); n == 0 && index < targetTaskIndex {
						// fmt.Printf("Swapping %s with %s", task.Name, targetTask.Name)
						// fmt.Println()
						// @TODO this is too long make it shorter
						sorter.Tasks[index], sorter.Tasks[targetTaskIndex] = sorter.Tasks[targetTaskIndex], sorter.Tasks[index]
					}
				}
			}
		}
	}
	return sorter.Tasks
}