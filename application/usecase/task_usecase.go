package usecase

import (
	"context"
	"sealion/domain/model"
	"sealion/domain/repository"
)

type TaskUseCase interface {
	GetTaskList(ctx context.Context) (*model.Tasks, error)
}

type taskUseCase struct {
	repository.TaskRepository
}

func NewTaskUseCase(r repository.TaskRepository) TaskUseCase {
	return &taskUseCase{r}
}

func (u *taskUseCase) GetTaskList(ctx context.Context) (*model.Tasks, error) {
	return nil, nil
}
