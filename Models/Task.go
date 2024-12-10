package Models

type Task struct {
	Name string `json:"name"`
	Command string `json:"command"`
	Requires []string `json:"requires"`
}