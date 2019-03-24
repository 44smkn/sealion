package repository

import (
	"context"
	"sealion/domain/model"
)

type TaskRepository interface {
	GetAll(ctx context.Context) ([]*model.Task, error)
	Add(ctx context.Context, task model.Task) (int64, error)
	Update(ctx context.Context, task model.Task) error
	Delete(ctx context.Context, id int64) error
}
