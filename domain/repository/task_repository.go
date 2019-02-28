package repository

import (
	"context"
	"sealion/domain/model"
)

type TaskRepository interface {
	getAll(ctx context.Context) (*model.Task, error)
}
