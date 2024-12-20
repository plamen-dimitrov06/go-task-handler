package tasksort

import(
	"encoding/json"
)

type JSONFormater struct {
}

/*
 * JSON implementation for the Formatter interface.
 */
func (formatter JSONFormater) Format(tasks []Task) string {
	var formattedTasks []JSONResponse
	for _, task := range tasks {
		outputTask := JSONResponse{Name: task.Name, Command: task.Command}
		formattedTasks = append(formattedTasks, outputTask)
	}
	t, _ := json.Marshal(formattedTasks)
	return string(t)
}
