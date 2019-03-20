package registry

import (
	"log"
	"sealion/application/usecase"
	"sealion/config"
	"sealion/domain/repository"
	"sealion/infrastructure/persistence/datastore"
	"sealion/interfaces/handler"
)

var Store *Container = inject()

func inject() *Container {

	conn, err := config.GetDbConn()
	if err != nil {
		log.Fatal("failed to new connection: ", err)
	}

	container := NewContainer()
	rd := &Definition{
		Name: "TaskRepository",
		Builder: func(c *Container) interface{} {
			return datastore.NewTaskRepository(conn)
		},
	}
	container.Register(rd)

	ud := &Definition{
		Name: "TaskUseCase",
		Builder: func(c *Container) interface{} {
			repo, _ := c.Get("TaskRepository").(repository.TaskRepository)
			return usecase.NewTaskUseCase(repo)
		},
	}
	container.Register(ud)

	hd := &Definition{
		Name: "AppHandler",
		Builder: func(c *Container) interface{} {
			uc, _ := c.Get("TaskUseCase").(usecase.TaskUseCase)
			return handler.NewTaskHandler(uc)
		},
	}
	container.Register(hd)

	return container
}
