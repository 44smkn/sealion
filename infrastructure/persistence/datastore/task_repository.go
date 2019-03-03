package datastore

import (
	"context"
	"database/sql"
	"sealion/domain/model"
	"sealion/domain/repository"

	_ "github.com/go-sql-driver/mysql"
)

type TaskRepository struct {
	Conn *sql.DB
}

func NewTaskRepository(conn *sql.DB) repository.TaskRepository {
	return &TaskRepository{Conn: conn}
}

func (r *TaskRepository) GetAll(ctx context.Context) ([]model.Task, error) {
	rows, err := r.Conn.Query("SELECT id, category, name, do_today, deadline FROM tasks")
	if err != nil {
		return nil, err
	}

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		err = rows.Scan(&t.Id, &t.Category, &t.Name, &t.DoToday, &t.Deadline)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}
