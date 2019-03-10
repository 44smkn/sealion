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

func (r *TaskRepository) GetAll(ctx context.Context) ([]*model.Task, error) {
	rows, err := r.Conn.Query("SELECT id, category, name, do_today, deadline FROM tasks")
	if err != nil {
		return nil, err
	}

	var tasks []*model.Task
	for rows.Next() {
		var t model.Task
		err = rows.Scan(&t.Id, &t.Category, &t.Name, &t.DoToday, &t.Deadline)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}
	return tasks, nil
}

func (r *TaskRepository) Add(ctx context.Context, task model.Task) (int64, error) {
	sql := `INSERT INTO tasks (category, name, do_today, deadline) VALUES (?, ?, ?, ?)`

	stmt, err := r.Conn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	row, err := stmt.Exec(task.Category, task.Name, task.DoToday, task.Deadline)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func (r *TaskRepository) Update(ctx context.Context, task model.Task) error {
	sql := `UPDATE tasks SET category=?, name=?, do_today=?, deadline=? WHERE id = ?`
	stmt, err := r.Conn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(task.Category, task.Name, task.DoToday, task.Deadline)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) Delete(ctx context.Context, id int64) error {
	sql := `DELETE FROM tasks WHERE id = ?`
	stmt, err := r.Conn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
