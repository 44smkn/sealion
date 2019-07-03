// +build wireinject

package main

import (
	"sealion/application/usecase"
	"sealion/config"
	"sealion/domain/service"
	"sealion/infrastructure/persistence/datastore"
	"sealion/interfaces/handler"

	"github.com/google/wire"
)

func initialize() (th handler.TaskHandler, err error) {

	wire.Build(
		config.Set,
		handler.Set,
		usecase.Set,
		service.Set,
		datastore.Set,
	)
	return
}
