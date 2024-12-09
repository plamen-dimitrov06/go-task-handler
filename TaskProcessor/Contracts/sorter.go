package TaskProcessor

import (
	"task-handler/TaskProcessor"
)

type Sorter interface {
	sort() []TaskProcessor.Task
}