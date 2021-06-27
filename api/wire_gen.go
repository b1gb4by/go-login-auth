// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"api/infrastructure/database"
	"api/interface/controller"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeControllers(db database.Connection) *controller.AppController {
	healthCheckController := controller.NewHealthCheckController(db)
	appController := controller.NewControllers(healthCheckController)
	return appController
}

// wire.go:

var controllerSet = wire.NewSet(controller.NewControllers, controller.NewHealthCheckController)

var interactorSet = wire.NewSet()

var gatewaySet = wire.NewSet()