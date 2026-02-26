package domains

import "time"

type Task struct {
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewTask(title string, description string) Task {
	return Task{
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}

type DomainGetTask struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewDomainGetTask(id int, title string, description string, completed bool, created_at time.Time, completed_at *time.Time) DomainGetTask {
	return DomainGetTask{
		ID:          id,
		Title:       title,
		Description: description,
		Completed:   completed,
		CreatedAt:   created_at,
		CompletedAt: completed_at,
	}
}

func (t *Task) Complete() {
	completeTime := time.Now()

	t.Completed = true
	t.CompletedAt = &completeTime
}

func (t *Task) Uncomplete() {
	t.Completed = false
	t.CompletedAt = nil
}
