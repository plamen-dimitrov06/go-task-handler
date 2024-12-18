package Contracts

import (
	"task-handler/Models"
)

type Formater interface {
	Format(tasks []Models.Task)
}
