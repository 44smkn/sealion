package usecase

import (
	"context"
	"log"
	"os"
	"sealion/domain/model"
	"sealion/domain/repository"
	"sealion/domain/service"
)

type TaskUseCase interface {
	GetTasks(ctx context.Context) ([]*model.Task, error)
	CreateTask(ctx context.Context, task model.Task) error
	UpdateTask(ctx context.Context, task model.Task) error
	DeleteTask(ctx context.Context, id int64) error
}

type taskUseCase struct {
	repository.TaskRepository
	service.TaskService
}

func NewTaskUseCase(r repository.TaskRepository, s service.TaskService) TaskUseCase {
	return &taskUseCase{r, s}
}

func (u *taskUseCase) GetTasks(ctx context.Context) ([]*model.Task, error) {
	if os.Getenv("SYNC_JIRA_ISSUE") == "on" {
		existedTasks, _ := u.GetTickets(ctx)
		jira, err := u.SyncJira(ctx, existedTasks)
		if err != nil {
			log.Println(err)
		}
		for _, j := range jira {
			_, err := u.Add(ctx, *j)
			if err != nil {
				return nil, err
			}
		}
	}
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

func (u *taskUseCase) UpdateTask(ctx context.Context, task model.Task) error {
	err := u.Update(ctx, task)
	if err != nil {
		return err
	}
	return nil
}

func (u *taskUseCase) DeleteTask(ctx context.Context, id int64) error {
	err := u.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
