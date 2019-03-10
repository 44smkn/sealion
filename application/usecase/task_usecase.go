package usecase

import (
	"context"
	"sealion/domain/model"
	"sealion/domain/repository"
	"sealion/infrastructure/persistence/datastore"
)

type TaskUseCase interface {
	GetTasks(ctx context.Context) ([]*model.Task, error)
}

type taskUseCase struct {
	repository.TaskRepository
}

func NewTaskUseCase(r repository.TaskRepository) TaskUseCase {
	return &taskUseCase{r}
}

func (u *taskUseCase) GetTasks(ctx context.Context) ([]*model.Task, error) {
	r := datastore.NewTaskRepository(u.Conn)
	tasks, err := r.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
