package tasksort

import (
	"encoding/json"
)

type JSONFormater struct {
}

/*
 * JSON implementation for the Formater interface.
 */
func (formater JSONFormater) Format(tasks []Task) string {
	var formatedTasks []JSONResponse
	for _, task := range tasks {
		outputTask := JSONResponse{Name: task.Name, Command: task.Command}
		formatedTasks = append(formatedTasks, outputTask)
	}
	t, _ := json.Marshal(formatedTasks)
	return string(t)
}

func NewJSONFormater() JSONFormater { return JSONFormater{} }
