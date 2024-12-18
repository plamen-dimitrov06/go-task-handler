# go-task-handler
Consume a list of tasks ([example](testdata/tasks.json)) through HTTP and order them while taking into consideration the "requires" property.
## Instructions
* docker compose up
* docker compose run go
* cd src/
* go run *.go
## Running tests
* docker compose run go
* cd src/TaskProcessor/
* go test
## TODO
* replace object instance with "new" factory functions
* make us of TableDrivenTest approach as opposed to AAA
* change naming of files to lowercase