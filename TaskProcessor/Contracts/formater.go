package TaskProcessor

import (
	"task-handler/TaskProcessor"
)

type Formater interface {
	format() []TaskProcessor.TaskOutputFormat
}