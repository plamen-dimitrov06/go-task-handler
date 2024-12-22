# go-task-handler
Consume a list of tasks ([example](testdata/tasks.json)) through HTTP and order them while taking into consideration the "requires" property.
## Instructions
* docker compose up
* docker compose run go
* cd src/cmd/tasksort
* go run *.go
## Running tests
* docker compose run go
* cd src/internal/tasksort
* go test