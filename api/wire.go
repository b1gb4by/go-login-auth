// +build wireinject

package main

import (
	"api/infrastructure/database"
	"api/interface/controller"
	"api/interface/gateway"
	"api/usecase/interactor"

	"github.com/google/wire"
)

var controllerSet = wire.NewSet(
	controller.NewControllers,
	controller.NewRegisterUserController,
	controller.NewHealthCheckController,
)

var interactorSet = wire.NewSet(
	interactor.NewRegisterUserInteractor,
)

var gatewaySet = wire.NewSet(
	gateway.NewRegisterUserGateway,
)

func InitializeControllers(db database.Connection, table string) *controller.AppController {
	wire.Build(controllerSet, interactorSet, gatewaySet)
	return nil
}
