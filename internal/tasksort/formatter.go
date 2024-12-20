package tasksort

type Formater interface {
	Format(tasks []Models.Task)
}
