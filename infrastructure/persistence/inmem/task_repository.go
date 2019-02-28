package inmem

import (
	"context"
	"sealion/domain/model"
	"sealion/domain/repository"
)

type TaskRepository struct{}

func NewTaskRepository() repository.TaskRepository {
	return &TaskRepository{}
}

func (r *TaskRepository) getAll(ctx context.Context) (*model.Task, error) {
	return nil, nil
}
