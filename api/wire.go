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
	controller.NewLoginAuthenticationController,
	controller.NewHealthCheckController,
)

var interactorSet = wire.NewSet(
	interactor.NewRegisterUserInteractor,
	interactor.NewLoginAuthenticationInteractor,
)

var gatewaySet = wire.NewSet(
	gateway.NewRegisterUserGateway,
	gateway.NewLoginAuthenticationGateway,
)

func InitializeControllers(db database.Connection, table string) *controller.AppController {
	wire.Build(controllerSet, interactorSet, gatewaySet)
	return nil
}
