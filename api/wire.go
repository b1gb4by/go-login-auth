// +build wireinject

package main

import (
	"api/config"
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
	controller.NewLogoutController,
	controller.NewUserAuthenticationController,
	controller.NewForgotController,
	controller.NewHealthCheckController,
)

var interactorSet = wire.NewSet(
	interactor.NewRegisterUserInteractor,
	interactor.NewLoginAuthenticationInteractor,
	interactor.NewUserAuthenticationInteractor,
	interactor.NewForgotInteractor,
)

var gatewaySet = wire.NewSet(
	gateway.NewRegisterUserGateway,
	gateway.NewLoginAuthenticationGateway,
	gateway.NewUserAuthenticationGateway,
	gateway.NewForgotGateway,
)

func InitializeControllers(db database.Connection, jc *config.JWTConfig) *controller.AppController {
	wire.Build(controllerSet, interactorSet, gatewaySet)
	return nil
}
