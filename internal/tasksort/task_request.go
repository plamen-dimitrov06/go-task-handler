package tasksort

type TaskRequest struct {
	Tasks []Task `json:"tasks"`
}

func NewTaskRequest() TaskRequest { return TaskRequest{} }