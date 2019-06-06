package datastore

import (
	"context"
	"database/sql"
	"sealion/domain/model"
	"sealion/domain/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

var Set = wire.NewSet(ProvideTaskRepository)

type TaskRepository struct {
	Conn *sql.DB
}

func ProvideTaskRepository(conn *sql.DB) repository.TaskRepository {
	return &TaskRepository{Conn: conn}
}

func (r *TaskRepository) GetAll(ctx context.Context) ([]*model.Task, error) {
	rows, err := r.Conn.Query("SELECT id, category, name, do_today, deadline, ticket_id, archive, description FROM Tasks")
	if err != nil {
		return nil, err
	}

	list, err := toTaskList(rows)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *TaskRepository) Add(ctx context.Context, task model.Task) (int64, error) {
	sql := `INSERT INTO Tasks (category, name, do_today, deadline, ticket_id, archive, description) VALUES (?, ?, ?, ?, ?, ?, ?)`

	stmt, err := r.Conn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	row, err := stmt.Exec(task.Category, task.Name, task.DoToday, task.Deadline, task.TicketId, task.Archive, task.Description)
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
	sql := `UPDATE Tasks SET category=?, name=?, do_today=?, deadline=?, ticket_id=?, archive=?, description=? WHERE id = ?`
	stmt, err := r.Conn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(task.Category, task.Name, task.DoToday, task.Deadline, task.TicketId, task.Archive, task.Description, task.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) Delete(ctx context.Context, id int64) error {
	sql := `DELETE FROM Tasks WHERE id = ?`
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

func (r *TaskRepository) GetTickets(ctx context.Context) ([]*model.Task, error) {
	rows, err := r.Conn.Query("SELECT id, category, name, do_today, deadline, ticket_id, archive, description FROM Tasks WHERE category = 'TICKET'")
	if err != nil {
		return nil, err
	}

	list, err := toTaskList(rows)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func toTaskList(rows *sql.Rows) ([]*model.Task, error) {
	tasks := []*model.Task{}
	for rows.Next() {
		var t model.Task
		err := rows.Scan(&t.Id, &t.Category, &t.Name, &t.DoToday, &t.Deadline, &t.TicketId, &t.Archive, &t.Description)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}
	return tasks, nil
}
