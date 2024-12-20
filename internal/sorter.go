package Contracts

import (
	"task-handler/Models"
)

type Sorter interface {
	Sort([]Models.Task) []Models.Task
}
