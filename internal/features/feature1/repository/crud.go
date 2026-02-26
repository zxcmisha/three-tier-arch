package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	ctx  context.Context
	conn *pgx.Conn
}

func NewPostgres(ctx context.Context, conn *pgx.Conn) *Postgres {
	return &Postgres{
		ctx:  ctx,
		conn: conn,
	}
}

func (p *Postgres) InsertTask(task TaskModel) error {
	sqlQuery := `
	INSERT INTO tasks (title, description, completed, created_at, completed_at)
	VALUES ($1, $2, $3, $4, $5);
	`
	if _, err := p.conn.Exec(p.ctx, sqlQuery, task.Title, task.Description, task.Completed, task.CreatedAt, task.CompletedAt); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) SelectTask(ID int) (TaskModel, error) {
	sqlQuery := `
	SELECT * FROM tasks
	WHERE id=$1
	ORDER BY id ASC;
	`
	row := p.conn.QueryRow(p.ctx, sqlQuery, ID)

	var task TaskModel
	if err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.CompletedAt); err != nil {
		return TaskModel{}, err
	}

	return task, nil
}

func (p *Postgres) SelectTasks() ([]TaskModel, error) {
	sqlQuery := `
	SELECT id, title, description, completed, created_at, completed_at FROM tasks
	ORDER BY id ASC;
	`
	tasks := make([]TaskModel, 0)
	rows, err := p.conn.Query(p.ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task TaskModel
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.CompletedAt); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (p *Postgres) SelectUncompletedTasks() ([]TaskModel, error) {
	sqlQuery := `
	SELECT * FROM tasks
	WHERE completed=FALSE
	ORDER BY id ASC;
	`
	tasks := make([]TaskModel, 0)
	rows, err := p.conn.Query(p.ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task TaskModel
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.CompletedAt); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (p *Postgres) CompleteTask(id int) error {
	sqlQuery := `
	UPDATE tasks
	SET completed=$1, completed_at=$2
	WHERE id=$3;
	`
	if _, err := p.conn.Exec(p.ctx, sqlQuery, true, time.Now(), id); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) UncompleteTask(id int) error {
	sqlQuery := `
	UPDATE tasks
	SET completed=$1, completed_at=$2
	WHERE id=$3;
	`
	if _, err := p.conn.Exec(p.ctx, sqlQuery, false, nil, id); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) DeleteTask(ID int) error {
	sqlQuery := `
	DELETE FROM tasks
	WHERE id=$1;
	`
	if _, err := p.conn.Exec(p.ctx, sqlQuery, ID); err != nil {
		return err
	}

	return nil
}
