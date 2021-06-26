package main

import (
	"api/config"
	"api/infrastructure/database"
	"api/infrastructure/http"
)

func main() {
	dbConf := config.NewDBConfig()
	// DB
	db := database.NewConnection(dbConf)
	defer db.Close()

	app := config.NewAppConfig()

	ctrls := InitializeControllers(db)

	routing := http.NewRouting(ctrls, app.Port)
	routing.SetRouting()
}
