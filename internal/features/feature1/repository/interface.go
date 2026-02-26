package repository

type Repository interface {
	InsertTask(task TaskModel) error
	SelectTask(ID int) (TaskModel, error)
	SelectTasks() ([]TaskModel, error)
	SelectUncompletedTasks() ([]TaskModel, error)
	CompleteTask(id int) error
	UncompleteTask(id int) error
	DeleteTask(ID int) error
}
