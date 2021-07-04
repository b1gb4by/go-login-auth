package main

import (
	"api/config"
	"api/infrastructure/database"
	"api/infrastructure/http"
)

func main() {
	dbConf := config.NewDBConfig()
	db := database.NewConnection(dbConf)
	defer db.Close()

	app := config.NewAppConfig()
	JWTConfig := config.NewJWTConfig()

	ctrls := InitializeControllers(db, dbConf.Table, JWTConfig)

	routing := http.NewRouting(ctrls, app.Port)
	routing.SetRouting()
}
