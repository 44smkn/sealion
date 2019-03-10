package repository

import (
	"context"
	"sealion/domain/model"
)

type TaskRepository interface {
	GetAll(ctx context.Context) ([]*model.Task, error)
}
