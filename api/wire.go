// +build wireinject

package main

import (
	"api/infrastructure/database"
	"api/interface/controller"

	"github.com/google/wire"
)

var controllerSet = wire.NewSet(
	controller.NewControllers,
	controller.NewHealthCheckController,
)

var interactorSet = wire.NewSet()

var gatewaySet = wire.NewSet()

func InitializeControllers(db database.Connection) *controller.AppController {
	wire.Build(controllerSet)
	return nil
}
