package registry

import (
	"log"
	"os"
	"sealion/application/usecase"
	"sealion/config"
	"sealion/domain/repository"
	"sealion/domain/service"
	"sealion/infrastructure/client"
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

	sd := &Definition{
		Name: "TaskService",
		Builder: func(c *Container) interface{} {
			baseUrl := os.Getenv("JIRA_BASE_URL")
			username := os.Getenv("JIRA_USERNAME")
			password := os.Getenv("JIRA_PASSWORD")
			client, _ := client.NewJira(baseUrl, username, password)
			return service.NewTaskService(client)
		},
	}
	container.Register(sd)

	ud := &Definition{
		Name: "TaskUseCase",
		Builder: func(c *Container) interface{} {
			repo, _ := c.Get("TaskRepository").(repository.TaskRepository)
			svc, _ := c.Get("TaskService").(service.TaskService)
			return usecase.NewTaskUseCase(repo, svc)
		},
	}
	container.Register(ud)

	h := &Definition{
		Name: "TaskHandler",
		Builder: func(c *Container) interface{} {
			uc, _ := c.Get("TaskUseCase").(usecase.TaskUseCase)
			return handler.NewTaskHandler(uc)
		},
	}
	container.Register(h)

	return container
}
