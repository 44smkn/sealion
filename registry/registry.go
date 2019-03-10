package registry

import (
	"net/http"
	"sealion/application/usecase"
	"sealion/domain/repository"
)

type Registry interface {
	NewTaskRepository() repository.TaskRepository
	NewTaskUseCase() usecase.TaskUseCase
	NewAppHandler() http.Handler
}

type registry struct {
	taskrepo    repository.TaskRepository
	taskuse     usecase.TaskUseCase
	taskHandler http.Handler
}

func NewRegistry() registry {

}

func (r *registry) NewTaskRepository() repository.TaskRepository {

}
