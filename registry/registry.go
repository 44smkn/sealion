package registry

import (
	"log"
	"net/http"
	"sealion/application/usecase"
	"sealion/config"
	"sealion/domain/repository"
	"sealion/infrastructure/persistence/datastore"
	"sealion/interfaces/handler"
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

func NewRegistry() *registry {
	return &registry{}
}

func (r *registry) NewTaskRepository() repository.TaskRepository {
	conn, err := config.GetDbConn()
	if err != nil {
		log.Fatal("failed to new connection: ", err)
	}
	return datastore.NewTaskRepository(conn)
}

func (r *registry) NewTaskUseCase() usecase.TaskUseCase {
	return usecase.NewTaskUseCase(r.taskrepo)
}

func (r *registry) NewAppHandler() http.Handler {
	return handler.NewTaskHandler(r.taskuse)
}
