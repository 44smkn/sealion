// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"sealion/application/usecase"
	"sealion/config"
	"sealion/domain/service"
	"sealion/infrastructure/client"
	"sealion/infrastructure/persistence/datastore"
	"sealion/interfaces/handler"
)

// Injectors from wire.go:

func initialize() (handler.TaskHandler, error) {
	db, err := config.GetDbConn()
	if err != nil {
		return nil, err
	}
	taskRepository := datastore.ProvideTaskRepository(db)
	jiraClient, err := client.NewJira()
	if err != nil {
		return nil, err
	}
	taskService := service.ProvideTaskService(jiraClient)
	taskUseCase := usecase.ProvideTaskUseCase(taskRepository, taskService)
	taskHandler := handler.ProvideTaskHandler(taskUseCase)
	return taskHandler, nil
}
