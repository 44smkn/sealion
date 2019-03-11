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
	r := &registry{}
	r.NewTaskRepository()
	r.NewTaskUseCase()
	r.NewAppHandler()
	return r
}

func (r *registry) NewTaskRepository() repository.TaskRepository {
	if r.taskrepo == nil {
		conn, err := config.GetDbConn()
		if err != nil {
			log.Fatal("failed to new connection: ", err)
		}
		r.taskrepo = datastore.NewTaskRepository(conn)
	}
	return r.taskrepo
}

func (r *registry) NewTaskUseCase() usecase.TaskUseCase {
	if r.taskuse == nil {
		r.taskuse = usecase.NewTaskUseCase(r.taskrepo)
	}
	return r.taskuse
}

func (r *registry) NewAppHandler() http.Handler {
	if r.taskHandler == nil {
		r.taskHandler = handler.NewTaskHandler(r.taskuse)
	}
	return r.taskHandler
}
