package tasksort

type Formater interface {
	Format(tasks []Task)
}
