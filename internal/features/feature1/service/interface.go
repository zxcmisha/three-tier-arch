package service

import (
	"restapi/internal/core/domains"
)

type Service interface {
	AddTask(task domains.Task) error
	GetTask(id int) (domains.DomainGetTask, error)
	GetAllTasks() ([]domains.DomainGetTask, error)
	GetAllUncompletedTasks() ([]domains.DomainGetTask, error)
	CompleteTask(id int) error
	UncompleteTask(id int) error
	DeleteTask(ID int) error
}
