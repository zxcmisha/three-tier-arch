package repository

import "time"

type TaskModel struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewTaskModel(title string, description string, completed bool, created_at time.Time, completed_at *time.Time) TaskModel {
	return TaskModel{
		Title:       title,
		Description: description,
		Completed:   completed,
		CreatedAt:   created_at,
		CompletedAt: completed_at,
	}
}
