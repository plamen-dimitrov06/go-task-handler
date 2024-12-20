package tasksort

/*
 * This struct is used for specifying the fields for the JSON response.
 */
type JSONResponse struct {
	Name    string `json:"name"`
	Command string `json:"command"`
}
