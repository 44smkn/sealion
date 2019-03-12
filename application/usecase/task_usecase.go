package usecase

import (
	"context"
	"sealion/domain/model"
	"sealion/domain/repository"
)

type TaskUseCase interface {
	GetTasks(ctx context.Context) ([]*model.Task, error)
	CreateTask(ctx context.Context, task model.Task) error
}

type taskUseCase struct {
	repository.TaskRepository
}

func NewTaskUseCase(r repository.TaskRepository) TaskUseCase {
	return &taskUseCase{r}
}

func (u *taskUseCase) GetTasks(ctx context.Context) ([]*model.Task, error) {
	tasks, err := u.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (u *taskUseCase) CreateTask(ctx context.Context, task model.Task) error {
	_, err := u.Add(ctx, task)
	if err != nil {
		return err
	}
	return nil
}
