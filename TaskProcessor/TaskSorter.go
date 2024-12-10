package TaskProcessor

import (
	"strings"
	"task-handler/Models"
)

type TaskSorter struct {
}

/**
 * Sort tasks taking into consideration the required property.
 */
func (sorter TaskSorter) Sort(tasks []Models.Task) []Models.Task {
	for index, task := range tasks {
		if len(task.Requires) > 0 {
			for _, requiredTask := range task.Requires {
				for targetIndex, targetTask := range tasks {
					if n := strings.Compare(requiredTask, targetTask.Name); n == 0 && index < targetIndex {
						// fmt.Printf("Swapping %s with %s", task.Name, targetTask.Name)
						// fmt.Println()
						tasks[index], tasks[targetIndex] = tasks[targetIndex], tasks[index]
					}
				}
			}
		}
	}
	return tasks
}