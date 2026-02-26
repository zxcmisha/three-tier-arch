package service

import (
	"restapi/internal/core/domains"
	"restapi/internal/features/feature1/repository"
	"sync"
)

type UserService struct {
	repository repository.Repository
	mtx        sync.RWMutex
}

func NewUserService(r repository.Repository) *UserService {
	return &UserService{
		repository: r,
	}
}

func (s *UserService) AddTask(task domains.Task) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	modelTask := repository.NewTaskModel(task.Title, task.Description, task.Completed, task.CreatedAt, task.CompletedAt)

	if err := s.repository.InsertTask(modelTask); err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetTask(id int) (domains.DomainGetTask, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	modelTask, err := s.repository.SelectTask(id)
	if err != nil {
		return domains.DomainGetTask{}, err
	}
	DomainTask := domains.NewDomainGetTask(modelTask.ID, modelTask.Title, modelTask.Description, modelTask.Completed, modelTask.CreatedAt, modelTask.CompletedAt)

	return DomainTask, nil
}

func (s *UserService) GetAllTasks() ([]domains.DomainGetTask, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	tasks := make([]domains.DomainGetTask, 0)
	taskmodels, err := s.repository.SelectTasks()
	if err != nil {
		return nil, err
	}

	for _, v := range taskmodels {
		DomainGetTask := domains.NewDomainGetTask(v.ID, v.Title, v.Description, v.Completed, v.CreatedAt, v.CompletedAt)
		tasks = append(tasks, DomainGetTask)
	}

	return tasks, nil
}

func (s *UserService) GetAllUncompletedTasks() ([]domains.DomainGetTask, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	UncompletedTasks := make([]domains.DomainGetTask, 0)
	taskmodels, err := s.repository.SelectUncompletedTasks()
	if err != nil {
		return nil, err
	}

	for _, v := range taskmodels {
		DomainGetTask := domains.NewDomainGetTask(v.ID, v.Title, v.Description, v.Completed, v.CreatedAt, v.CompletedAt)
		UncompletedTasks = append(UncompletedTasks, DomainGetTask)
	}

	return UncompletedTasks, nil
}

func (s *UserService) CompleteTask(id int) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if err := s.repository.CompleteTask(id); err != nil {
		return err
	}

	return nil
}

func (s *UserService) UncompleteTask(id int) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if err := s.repository.UncompleteTask(id); err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteTask(id int) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if err := s.repository.DeleteTask(id); err != nil {
		return err
	}

	return nil
}
